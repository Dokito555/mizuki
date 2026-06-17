package services

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sync"

	"github.com/Dokito555/mizuki/internal/constants"
	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/Dokito555/mizuki/internal/services/detection"
	"github.com/Dokito555/mizuki/internal/services/pcap"
	"github.com/sirupsen/logrus"
)

type UploadService struct {
	uploadRepo      repositories.UploadRepository
	flowRepo        repositories.FlowRepository
	pcapEngine      *pcap.Engine
	detectionEngine *detection.DetectionEngine
	log             *logrus.Logger
	maxFileSize     int64
	mu              sync.Mutex
	cancelFuncs     map[uint]context.CancelFunc
}

func NewUploadService(
	uploadRepo repositories.UploadRepository,
	flowRepo repositories.FlowRepository,
	pcapEngine *pcap.Engine,
	detectionEngine *detection.DetectionEngine,
	log *logrus.Logger,
	maxFileSize int64,
) *UploadService {
	return &UploadService{
		uploadRepo:      uploadRepo,
		flowRepo:        flowRepo,
		pcapEngine:      pcapEngine,
		detectionEngine: detectionEngine,
		log:             log,
		maxFileSize:     maxFileSize,
		cancelFuncs:     make(map[uint]context.CancelFunc),
	}
}

func (s *UploadService) ProcessUpload(ctx context.Context, file multipart.File, header *multipart.FileHeader, forceReparse bool) (*models.UploadResponse, error) {
	if header.Size > s.maxFileSize {
		return nil, fmt.Errorf("uploadService.ProcessUpload: %w", constants.ErrFileTooLarge)
	}

	tempDir := os.Getenv("MIZUKI_TEMP_DIR")
	tmpFile, err := os.CreateTemp(tempDir, "mizuki-*.pcap")
	if err != nil {
		return nil, fmt.Errorf("uploadService.ProcessUpload temp: %w", err)
	}

	// copy + hash must happen here, before the multipart file is closed
	hash := sha256.New()
	written, err := io.Copy(tmpFile, io.TeeReader(file, hash))
	if err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("uploadService.ProcessUpload copy: %w", err)
	}

	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("uploadService.ProcessUpload seek: %w", err)
	}

	fileHash := fmt.Sprintf("%x", hash.Sum(nil))

	if !forceReparse {
		existing, err := s.uploadRepo.FindByHash(ctx, fileHash)
		if err == nil && existing.Status == entities.UploadDone {
			tmpFile.Close()
			os.Remove(tmpFile.Name())
			resp := toUploadResponse(existing)
			return &resp, nil
		}
	}

	upload := &entities.Upload{
		Filename: header.Filename,
		FileSize: written,
		FileHash: fileHash,
		Status:   entities.UploadQueued,
	}

	if err := s.uploadRepo.Create(ctx, upload); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("uploadService.ProcessUpload create: %w", err)
	}

	s.mu.Lock()
	pCtx, cancel := context.WithCancel(context.Background())
	s.cancelFuncs[upload.ID] = cancel
	s.mu.Unlock()

	go s.runPipeline(pCtx, upload.ID, tmpFile)

	resp := toUploadResponse(upload)
	return &resp, nil
}

func (s *UploadService) CancelUpload(uploadID uint) bool {
	s.mu.Lock()
	cancel, ok := s.cancelFuncs[uploadID]
	if ok {
		delete(s.cancelFuncs, uploadID)
	}
	s.mu.Unlock()
	if ok {
		cancel()
		return true
	}
	return false
}

func (s *UploadService) runPipeline(ctx context.Context, uploadID uint, file *os.File) {
	tmpPath := file.Name()

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		s.log.WithField("upload_id", uploadID).Errorf("pipeline panic: %v", r)
	// 		upload, err := s.uploadRepo.FindByID(ctx, uploadID)
	// 		if err == nil {
	// 			upload.Status = entities.UploadError
	// 			upload.ErrorMsg = fmt.Sprintf("panic: %v", r)
	// 			upload.ProgressPct = 100
	// 			s.uploadRepo.Update(ctx, upload)
	// 		}
	// 	}
	// }()

	defer file.Close()
	defer func() {
		if err := os.Remove(tmpPath); err != nil && !os.IsNotExist(err) {
			s.log.WithField("upload_id", uploadID).Errorf("failed to remove temp file %s: %v", tmpPath, err)
		}
	}()
	defer func() {
		s.mu.Lock()
		delete(s.cancelFuncs, uploadID)
		s.mu.Unlock()
	}()

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
		SamplePacketLen:    100,
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

	go func() {
		if err := s.detectionEngine.AnalyzeUpload(context.Background(), uploadID); err != nil {
			log.Errorf("detection analysis failed: %v", err)
		}
	}()
}

func (s *UploadService) AnalyzeUpload(ctx context.Context, uploadID uint) error {
	return s.detectionEngine.AnalyzeUpload(ctx, uploadID)
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
