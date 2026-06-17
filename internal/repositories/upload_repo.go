package repositories

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/constants"
	"github.com/Dokito555/mizuki/internal/entities"
	"gorm.io/gorm"
)

type UploadRepository interface {
	Create(ctx context.Context, upload *entities.Upload) error
	FindByID(ctx context.Context, id uint) (*entities.Upload, error)
	FindByHash(ctx context.Context, hash string) (*entities.Upload, error)
	Update(ctx context.Context, upload *entities.Upload) error
	UpdateProgress(ctx context.Context, id uint, status entities.UploadStatus, pct int, packetsProcessed int64) error
	List(ctx context.Context, page, pageSize int) ([]entities.Upload, int64, error)
}

type uploadRepository struct {
	Repository[entities.Upload]
	db *gorm.DB
}

func NewUploadRepository(db *gorm.DB) UploadRepository {
	return &uploadRepository{
		Repository: Repository[entities.Upload]{DB: db},
		db:         db,
	}
}

func (r *uploadRepository) Create(ctx context.Context, upload *entities.Upload) error {
	return r.Repository.Create(ctx, upload)
}

func (r *uploadRepository) FindByID(ctx context.Context, id uint) (*entities.Upload, error) {
	var upload entities.Upload
	if err := r.db.WithContext(ctx).First(&upload, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("uploadRepo.FindByID(%d): %w", id, constants.ErrUploadNotFound)
		}
		return nil, fmt.Errorf("uploadRepo.FindByID(%d): %w", id, err)
	}
	return &upload, nil
}

func (r *uploadRepository) FindByHash(ctx context.Context, hash string) (*entities.Upload, error) {
	var upload entities.Upload
	if err := r.db.WithContext(ctx).Where("file_hash = ?", hash).Order("created_at DESC").First(&upload).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("uploadRepo.FindByHash: %w", constants.ErrUploadNotFound)
		}
		return nil, fmt.Errorf("uploadRepo.FindByHash: %w", err)
	}
	return &upload, nil
}

func (r *uploadRepository) Update(ctx context.Context, upload *entities.Upload) error {
	return r.Repository.Update(ctx, upload)
}

func (r *uploadRepository) UpdateProgress(ctx context.Context, id uint, status entities.UploadStatus, pct int, packetsProcessed int64) error {
	result := r.db.WithContext(ctx).Model(&entities.Upload{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":            status,
		"progress_pct":      pct,
		"packets_processed": packetsProcessed,
	})
	if result.Error != nil {
		return fmt.Errorf("uploadRepo.UpdateProgress(%d): %w", id, result.Error)
	}
	return nil
}

func (r *uploadRepository) List(ctx context.Context, page, pageSize int) ([]entities.Upload, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	var total int64
	if err := r.db.WithContext(ctx).Model(&entities.Upload{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("uploadRepo.List.count: %w", err)
	}

	var uploads []entities.Upload
	offset := (page - 1) * pageSize
	if err := r.db.WithContext(ctx).Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&uploads).Error; err != nil {
		return nil, 0, fmt.Errorf("uploadRepo.List: %w", err)
	}

	return uploads, total, nil
}
