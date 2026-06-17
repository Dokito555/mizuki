package repositories

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/constants"
	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/Dokito555/mizuki/internal/models"
	"gorm.io/gorm"
)

type FlowRepository interface {
	Create(ctx context.Context, flow *entities.Flow) error
	CreateBatch(ctx context.Context, flows []entities.Flow) error
	CreatePacketSamples(ctx context.Context, samples []entities.FlowPacketSample) error
	FindByID(ctx context.Context, id uint) (*entities.Flow, error)
	FindAll(ctx context.Context, filter models.FlowFilter) ([]entities.Flow, int64, error)
	FindPacketSamplesByFlowID(ctx context.Context, flowID uint, limit int) ([]entities.FlowPacketSample, error)
	UpdateScores(ctx context.Context, flows []entities.Flow) error
	DeleteByUploadID(ctx context.Context, uploadID uint) error
	CountByUploadID(ctx context.Context, uploadID uint) (int64, error)
}

type flowRepository struct {
	Repository[entities.Flow]
	db *gorm.DB
}

func NewFlowRepository(db *gorm.DB) FlowRepository {
	return &flowRepository{
		Repository: Repository[entities.Flow]{DB: db},
		db:         db,
	}
}

func (r *flowRepository) Create(ctx context.Context, flow *entities.Flow) error {
	return r.Repository.Create(ctx, flow)
}

func (r *flowRepository) CreateBatch(ctx context.Context, flows []entities.Flow) error {
	if len(flows) == 0 {
		return nil
	}
	batchSize := 1000
	if err := r.db.WithContext(ctx).CreateInBatches(flows, batchSize).Error; err != nil {
		return fmt.Errorf("flowRepo.CreateBatch: %w", err)
	}
	return nil
}

func (r *flowRepository) CreatePacketSamples(ctx context.Context, samples []entities.FlowPacketSample) error {
	if len(samples) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).CreateInBatches(samples, 1000).Error; err != nil {
		return fmt.Errorf("flowRepo.CreatePacketSamples: %w", err)
	}
	return nil
}

func (r *flowRepository) FindByID(ctx context.Context, id uint) (*entities.Flow, error) {
	var flow entities.Flow
	if err := r.db.WithContext(ctx).First(&flow, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("flowRepo.FindByID(%d): %w", id, constants.ErrFlowNotFound)
		}
		return nil, fmt.Errorf("flowRepo.FindByID(%d): %w", id, err)
	}
	return &flow, nil
}

func (r *flowRepository) FindAll(ctx context.Context, filter models.FlowFilter) ([]entities.Flow, int64, error) {
	filter.Normalize()
	query := r.db.WithContext(ctx).Model(&entities.Flow{})

	if filter.SrcIP != "" {
		query = query.Where("src_ip = ?", filter.SrcIP)
	}
	if filter.DstIP != "" {
		query = query.Where("dst_ip = ?", filter.DstIP)
	}
	if filter.Protocol != "" {
		query = query.Where("protocol = ?", filter.Protocol)
	}
	if filter.MinScore > 0 {
		query = query.Where("score >= ?", filter.MinScore)
	}
	if filter.Since != nil {
		query = query.Where("first_seen >= ?", filter.Since)
	}
	if filter.Until != nil {
		query = query.Where("last_seen <= ?", filter.Until)
	}
	if filter.UploadID > 0 {
		query = query.Where("raw_file_id = ?", filter.UploadID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("flowRepo.FindAll.count: %w", err)
	}

	sortOrder := filter.SortBy
	if filter.SortDesc {
		sortOrder += " DESC"
	}

	var flows []entities.Flow
	offset := (filter.Page - 1) * filter.PageSize
	if err := query.Order(sortOrder).Limit(filter.PageSize).Offset(offset).Find(&flows).Error; err != nil {
		return nil, 0, fmt.Errorf("flowRepo.FindAll: %w", err)
	}

	return flows, total, nil
}

func (r *flowRepository) FindPacketSamplesByFlowID(ctx context.Context, flowID uint, limit int) ([]entities.FlowPacketSample, error) {
	var samples []entities.FlowPacketSample
	if err := r.db.WithContext(ctx).Where("flow_id = ?", flowID).Order("timestamp ASC").Limit(limit).Find(&samples).Error; err != nil {
		return nil, fmt.Errorf("flowRepo.FindPacketSamplesByFlowID: %w", err)
	}
	return samples, nil
}

func (r *flowRepository) UpdateScores(ctx context.Context, flows []entities.Flow) error {
	if len(flows) == 0 {
		return nil
	}
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for i := range flows {
		if err := tx.Model(&flows[i]).Updates(map[string]interface{}{
			"score":   flows[i].Score,
			"threats": flows[i].Threats,
		}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("flowRepo.UpdateScores(%d): %w", flows[i].ID, err)
		}
	}
	return tx.Commit().Error
}

func (r *flowRepository) DeleteByUploadID(ctx context.Context, uploadID uint) error {
	if err := r.db.WithContext(ctx).Where("raw_file_id = ?", uploadID).Delete(&entities.Flow{}).Error; err != nil {
		return fmt.Errorf("flowRepo.DeleteByUploadID(%d): %w", uploadID, err)
	}
	return nil
}

func (r *flowRepository) CountByUploadID(ctx context.Context, uploadID uint) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entities.Flow{}).Where("raw_file_id = ?", uploadID).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("flowRepo.CountByUploadID(%d): %w", uploadID, err)
	}
	return count, nil
}
