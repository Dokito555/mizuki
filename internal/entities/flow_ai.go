package entities

import (
	"encoding/json"
	"time"
)

type FlowAI struct {
	ID        uint             `gorm:"primaryKey"`
	FlowID    uint             `gorm:"not null;uniqueIndex;constraint:OnDelete:CASCADE"`
	Analysis  *json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
