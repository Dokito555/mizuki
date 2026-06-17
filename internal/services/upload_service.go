package services

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Dokito555/mizuki/internal/constants"
	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/Dokito555/mizuki/internal/services/pcap"
	"github.com/sirupsen/logrus"
)

type UploadService struct {
	uploadRepo  repositories.UploadRepository
	flowRepo    repositories.FlowRepository
	pcapEngine  *pcap.Engine
	flowSvc     *FlowService
	log         *logrus.Logger
	maxFileSize int64
}

func NewUploadService(
	uploadRepo repositories.UploadRepository,
	flowRepo repositories.FlowRepository,
	pcapEngine *pcap.Engine,
	flowSvc *FlowService,
	log *logrus.Logger,
	maxFileSize int64,
) *UploadService {
	return &UploadService{
		uploadRepo:  uploadRepo,
		flowRepo:    flowRepo,
		pcapEngine:  pcapEngine,
		flowSvc:     flowSvc,
		log:         log,
		maxFileSize: maxFileSize,
	}
}

func (s *UploadService) ProcessUpload(ctx context.Context, file multipart.File, header *multipart.FileHeader, forceReparse bool) (*models.UploadResponse, error) {
	if header.Size > s.maxFileSize {
		return nil, fmt.Errorf("uploadService.ProcessUpload: %w", constants.ErrFileTooLarge)
	}

	fileHash, err := hashReader(file)
	if err != nil {
		return nil, fmt.Errorf("uploadService.ProcessUpload hash: %w", err)
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("uploadService.ProcessUpload seek: %w", err)
	}

	if !forceReparse {
		existing, err := s.uploadRepo.FindByHash(ctx, fileHash)
		if err == nil && existing.Status == entities.UploadDone {
			resp := toUploadResponse(existing)
			return &resp, nil
		}
	}

	upload := &entities.Upload{
		Filename: header.Filename,
		FileSize: header.Size,
		FileHash: fileHash,
		Status:   entities.UploadQueued,
	}

	if err := s.uploadRepo.Create(ctx, upload); err != nil {
		return nil, fmt.Errorf("uploadService.ProcessUpload create: %w", err)
	}

	go s.runPipeline(context.Background(), upload.ID, file, header.Filename)

	resp := toUploadResponse(upload)
	return &resp, nil
}

func (s *UploadService) runPipeline(ctx context.Context, uploadID uint, file multipart.File, filename string) {
	log := s.log.WithField("upload_id", uploadID)

	updateStatus := func(status entities.UploadStatus, pct int, packets int64) {
		if err := s.uploadRepo.UpdateProgress(ctx, uploadID, status, pct, packets); err != nil {
			log.Errorf("update progress: %v", err)
		}
	}

	markError := func(msg string) {
		upload, err := s.uploadRepo.FindByID(ctx, uploadID)
		if err != nil {
			log.Errorf("find upload: %v", err)
			return
		}
		upload.Status = entities.UploadError
		upload.ErrorMsg = msg
		upload.ProgressPct = 100
		if err := s.uploadRepo.Update(ctx, upload); err != nil {
			log.Errorf("mark error: %v", err)
		}
	}

	updateStatus(entities.UploadParsing, 0, 0)

	result, err := s.pcapEngine.Parse(ctx, file, pcap.ParseParams{
		MergeBidirectional: false,
		SamplePayloadLen:   64,
		OnProgress: func(packets int64, pct int) {
			updateStatus(entities.UploadParsing, 0, packets)
		},
	})
	if err != nil {
		markError(fmt.Sprintf("parse failed: %v", err))
		return
	}

	updateStatus(entities.UploadInserting, 50, result.TotalPackets)

	flows := make([]entities.Flow, 0, len(result.FlowStats))
	for _, stats := range result.FlowStats {
		flow := pcap.ToFlowEntity(stats, uploadID, nil)
		flows = append(flows, *flow)
	}

	if err := s.flowRepo.CreateBatch(ctx, flows); err != nil {
		markError(fmt.Sprintf("insert flows failed: %v", err))
		return
	}

	var allSamples []entities.FlowPacketSample
	for _, flow := range flows {
		for _, stats := range result.FlowStats {
			if stats.Key.String() == fmt.Sprintf("%s:%d-%s:%d-%s",
				flow.SrcIP, flow.SrcPort, flow.DstIP, flow.DstPort, flow.Protocol) {
				samples := pcap.ToPacketSampleEntities(flow.ID, stats)
				if samples != nil {
					allSamples = append(allSamples, samples...)
				}
				break
			}
		}
	}

	if len(allSamples) > 0 {
		if err := s.flowRepo.CreatePacketSamples(ctx, allSamples); err != nil {
			log.Errorf("insert packet samples: %v", err)
		}
	}

	upload, err := s.uploadRepo.FindByID(ctx, uploadID)
	if err != nil {
		log.Errorf("find upload after insert: %v", err)
		return
	}
	upload.Status = entities.UploadDone
	upload.ProgressPct = 100
	upload.FlowCount = len(flows)
	upload.PacketsProcessed = result.TotalPackets
	upload.DurationMs = result.Duration.Seconds() * 1000
	if err := s.uploadRepo.Update(ctx, upload); err != nil {
		log.Errorf("finalize upload: %v", err)
	}

	log.WithFields(logrus.Fields{
		"flows":   len(flows),
		"packets": result.TotalPackets,
		"elapsed": result.Duration.String(),
	}).Info("upload processing complete")
}

func (s *UploadService) GetUploadByID(ctx context.Context, id uint) (*models.UploadResponse, error) {
	upload, err := s.uploadRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("uploadService.GetUploadByID(%d): %w", id, err)
	}
	resp := toUploadResponse(upload)
	return &resp, nil
}

func (s *UploadService) ListUploads(ctx context.Context, page, pageSize int) (*models.PaginatedResponse, error) {
	uploads, total, err := s.uploadRepo.List(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("uploadService.ListUploads: %w", err)
	}

	items := make([]models.UploadResponse, len(uploads))
	for i, u := range uploads {
		items[i] = toUploadResponse(&u)
	}

	resp := models.NewPaginated(items, page, pageSize, total)
	return &resp, nil
}

func (s *UploadService) Reparse(ctx context.Context, uploadID uint) (*models.UploadResponse, error) {
	upload, err := s.uploadRepo.FindByID(ctx, uploadID)
	if err != nil {
		return nil, fmt.Errorf("uploadService.Reparse(%d): %w", uploadID, err)
	}

	if err := s.flowRepo.DeleteByUploadID(ctx, uploadID); err != nil {
		return nil, fmt.Errorf("uploadService.Reparse delete flows: %w", err)
	}

	upload.Status = entities.UploadQueued
	upload.ProgressPct = 0
	upload.PacketsProcessed = 0
	upload.FlowCount = 0
	upload.DurationMs = 0
	upload.ErrorMsg = ""
	if err := s.uploadRepo.Update(ctx, upload); err != nil {
		return nil, fmt.Errorf("uploadService.Reparse reset: %w", err)
	}

	resp := toUploadResponse(upload)
	return &resp, nil
}

func hashReader(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func toUploadResponse(u *entities.Upload) models.UploadResponse {
	return models.UploadResponse{
		ID:               u.ID,
		Filename:         u.Filename,
		FileSize:         u.FileSize,
		FileHash:         u.FileHash,
		FileType:         u.FileType,
		Status:           string(u.Status),
		ProgressPct:      u.ProgressPct,
		PacketsProcessed: u.PacketsProcessed,
		FlowCount:        u.FlowCount,
		DurationMs:       u.DurationMs,
		ErrorMsg:         u.ErrorMsg,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}
