package entities

import (
	"time"
)

type UploadStatus string

const (
	UploadQueued    UploadStatus = "queued"
	UploadHashing   UploadStatus = "hashing"
	UploadParsing   UploadStatus = "parsing"
	UploadInserting UploadStatus = "inserting"
	UploadDone      UploadStatus = "done"
	UploadError     UploadStatus = "error"
)

type Upload struct {
	ID               uint         `gorm:"primaryKey"`
	Filename         string       `gorm:"type:varchar(255);not null"`
	FileSize         int64        `gorm:"not null"`
	FileHash         string       `gorm:"type:varchar(64);not null;index"`
	FileType         string       `gorm:"type:varchar(16)"`
	Status           UploadStatus `gorm:"type:varchar(16);not null;default:'queued';index"`
	ProgressPct      int          `gorm:"not null;default:0"`
	PacketsProcessed int64        `gorm:"not null;default:0"`
	FlowCount        int          `gorm:"not null;default:0"`
	DurationMs       float64
	ErrorMsg         string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
