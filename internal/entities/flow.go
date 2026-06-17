package entities

import (
	"time"
)

type Flow struct {
	ID            uint      `gorm:"primaryKey"`
	SrcIP         string    `gorm:"type:inet;not null"`
	DstIP         string    `gorm:"type:inet;not null"`
	SrcPort       int       `gorm:"not null"`
	DstPort       int       `gorm:"not null"`
	Protocol      string    `gorm:"type:varchar(8);not null;index"`
	FirstSeen     time.Time `gorm:"not null;index"`
	LastSeen      time.Time `gorm:"not null"`
	PacketCount   int64     `gorm:"not null;default:0"`
	ByteCount     int64     `gorm:"not null;default:0"`
	SrcMAC        string    `gorm:"type:macaddr"`
	DstMAC        string    `gorm:"type:macaddr"`
	TLSVersion    string    `gorm:"type:varchar(16)"`
	TLSSNI        string    `gorm:"type:varchar(255)"`
	DNSQueries    []string  `gorm:"type:text[]"`
	AppProtocol   string    `gorm:"type:varchar(32)"`
	PayloadSample []byte    `gorm:"type:bytea"`

	IATAvgMs    float64
	IATMinMs    float64
	IATMaxMs    float64
	IATStdDevMs float64

	Score     float64 `gorm:"default:0;index"`
	RawFileID uint    `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FlowPacketSample struct {
	ID        uint      `gorm:"primaryKey"`
	FlowID    uint      `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	Timestamp time.Time `gorm:"not null;index"`
	Size      int       `gorm:"not null"`
}
