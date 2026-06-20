package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Dokito555/mizuki/internal/entities"
	"gorm.io/gorm"
)

type UploadAIBatchRepository interface {
	FindByUploadID(ctx context.Context, uploadID uint) (*entities.UploadAIBatch, error)
	CreateOrUpdate(ctx context.Context, batch *entities.UploadAIBatch) error
}

type uploadAIBatchRepository struct {
	db *gorm.DB
}

func NewUploadAIBatchRepository(db *gorm.DB) UploadAIBatchRepository {
	return &uploadAIBatchRepository{db: db}
}

func (r *uploadAIBatchRepository) FindByUploadID(ctx context.Context, uploadID uint) (*entities.UploadAIBatch, error) {
	var batch entities.UploadAIBatch
	if err := r.db.WithContext(ctx).Where("upload_id = ?", uploadID).First(&batch).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("uploadAIBatchRepo.FindByUploadID(%d): %w", uploadID, err)
	}
	return &batch, nil
}

func (r *uploadAIBatchRepository) CreateOrUpdate(ctx context.Context, batch *entities.UploadAIBatch) error {
	if batch.ID == 0 {
		batch.CreatedAt = time.Now()
	}
	batch.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(batch).Error
}
