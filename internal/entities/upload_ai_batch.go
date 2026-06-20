package entities

import "time"

type UploadAIBatch struct {
	ID              uint       `gorm:"primaryKey"`
	UploadID        uint       `gorm:"not null;uniqueIndex;constraint:OnDelete:CASCADE"`
	Status          string     `gorm:"type:varchar(16);not null;default:'pending'"`
	ProcessedCount  int        `gorm:"not null;default:0"`
	TotalCount      int        `gorm:"not null;default:0"`
	ErrorCount      int        `gorm:"not null;default:0"`
	StartedAt       *time.Time
	CompletedAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
