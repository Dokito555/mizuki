package repositories

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/entities"
	"gorm.io/gorm"
)

type FlowAIRepository interface {
	Save(ctx context.Context, flowID uint, analysis []byte) error
	FindByFlowID(ctx context.Context, flowID uint) (*entities.FlowAI, error)
	DeleteByFlowID(ctx context.Context, flowID uint) error
}

type flowAIRepository struct {
	db *gorm.DB
}

func NewFlowAIRepository(db *gorm.DB) FlowAIRepository {
	return &flowAIRepository{db: db}
}

func (r *flowAIRepository) Save(ctx context.Context, flowID uint, analysis []byte) error {
	result := r.db.WithContext(ctx).Exec(
		`INSERT INTO flow_ais (flow_id, analysis, created_at, updated_at)
		 VALUES (?, ?, NOW(), NOW())
		 ON CONFLICT (flow_id) DO UPDATE SET analysis = ?, updated_at = NOW()`,
		flowID, analysis, analysis,
	)
	if result.Error != nil {
		return fmt.Errorf("flowAIRepo.Save(%d): %w", flowID, result.Error)
	}
	return nil
}

func (r *flowAIRepository) FindByFlowID(ctx context.Context, flowID uint) (*entities.FlowAI, error) {
	var flowAI entities.FlowAI
	if err := r.db.WithContext(ctx).Where("flow_id = ?", flowID).First(&flowAI).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("flowAIRepo.FindByFlowID(%d): %w", flowID, err)
	}
	return &flowAI, nil
}

func (r *flowAIRepository) DeleteByFlowID(ctx context.Context, flowID uint) error {
	if err := r.db.WithContext(ctx).Where("flow_id = ?", flowID).Delete(&entities.FlowAI{}).Error; err != nil {
		return fmt.Errorf("flowAIRepo.DeleteByFlowID(%d): %w", flowID, err)
	}
	return nil
}
