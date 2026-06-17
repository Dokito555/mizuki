package models

import "time"

type UploadResponse struct {
	ID               uint      `json:"id"`
	Filename         string    `json:"filename"`
	FileSize         int64     `json:"file_size"`
	FileHash         string    `json:"file_hash"`
	FileType         string    `json:"file_type"`
	Status           string    `json:"status"`
	ProgressPct      int       `json:"progress_pct"`
	PacketsProcessed int64     `json:"packets_processed"`
	FlowCount        int       `json:"flow_count"`
	DurationMs       float64   `json:"duration_ms"`
	ErrorMsg         string    `json:"error,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
