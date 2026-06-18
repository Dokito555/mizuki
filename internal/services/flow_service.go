package services

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/sirupsen/logrus"
)

type FlowService struct {
	flowRepo   repositories.FlowRepository
	uploadRepo repositories.UploadRepository
	log        *logrus.Logger
}

func NewFlowService(
	flowRepo repositories.FlowRepository,
	uploadRepo repositories.UploadRepository,
	log *logrus.Logger,
) *FlowService {
	return &FlowService{
		flowRepo:   flowRepo,
		uploadRepo: uploadRepo,
		log:        log,
	}
}

func (s *FlowService) GetByID(ctx context.Context, id uint) (*models.FlowDetail, error) {
	flow, err := s.flowRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("flowService.GetByID(%d): %w", id, err)
	}

	detail := &models.FlowDetail{
		FlowResponse: toFlowResponse(flow),
	}

	samples, err := s.flowRepo.FindPacketSamplesByFlowID(ctx, id, 100)
	if err != nil {
		s.log.Errorf("flowService.GetByID(%d) packet samples: %v", id, err)
	}
	if len(samples) > 0 {
		detail.PacketSamples = make([]models.PacketSampleItem, len(samples))
		for i, s := range samples {
			detail.PacketSamples[i] = models.PacketSampleItem{
				Timestamp: s.Timestamp,
				Size:      s.Size,
			}
		}
	}

	return detail, nil
}

func (s *FlowService) List(ctx context.Context, filter models.FlowFilter) (*models.PaginatedResponse, error) {
	flows, total, err := s.flowRepo.FindAll(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("flowService.List: %w", err)
	}

	items := make([]models.FlowResponse, len(flows))
	for i, f := range flows {
		items[i] = toFlowResponse(&f)
	}

	resp := models.NewPaginated(items, filter.Page, filter.PageSize, total)
	return &resp, nil
}

func toFlowResponse(f *entities.Flow) models.FlowResponse {
	return models.FlowResponse{
		ID:          f.ID,
		SrcIP:       f.SrcIP,
		DstIP:       f.DstIP,
		SrcPort:     f.SrcPort,
		DstPort:     f.DstPort,
		Protocol:    f.Protocol,
		FirstSeen:   f.FirstSeen,
		LastSeen:    f.LastSeen,
		PacketCount: f.PacketCount,
		ByteCount:   f.ByteCount,
		SrcMAC:      f.SrcMAC,
		DstMAC:      f.DstMAC,
		TLSVersion:  f.TLSVersion,
		TLSSNI:      f.TLSSNI,
		DNSQueries:  []string(f.DNSQueries),
		AppProtocol: f.AppProtocol,
		IATAvgMs:    f.IATAvgMs,
		IATMinMs:    f.IATMinMs,
		IATMaxMs:    f.IATMaxMs,
		IATStdDevMs: f.IATStdDevMs,
		Score:       f.Score,
		Threats:     []string(f.Threats),
		CreatedAt:   f.CreatedAt,
	}
}

func (s *FlowService) DeleteByUploadID(ctx context.Context, uploadID uint) error {
	return s.flowRepo.DeleteByUploadID(ctx, uploadID)
}
