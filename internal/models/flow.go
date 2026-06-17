package models

import "time"

type FlowFilter struct {
	SrcIP    string     `form:"src_ip"`
	DstIP    string     `form:"dst_ip"`
	Protocol string     `form:"protocol"`
	MinScore float64    `form:"min_score"`
	Since    *time.Time `form:"since"`
	Until    *time.Time `form:"until"`
	UploadID uint       `form:"upload_id"`
	Page     int        `form:"page"`
	PageSize int        `form:"page_size"`
	SortBy   string     `form:"sort_by"`
	SortDesc bool       `form:"sort_desc"`
}

const (
	DefaultPageSize = 20
	MaxPageSize     = 500
)

func (f *FlowFilter) Normalize() {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.PageSize < 1 {
		f.PageSize = DefaultPageSize
	}
	if f.PageSize > MaxPageSize {
		f.PageSize = MaxPageSize
	}
	if f.SortBy == "" {
		f.SortBy = "first_seen"
		f.SortDesc = true
	}
}

type FlowResponse struct {
	ID          uint      `json:"id"`
	SrcIP       string    `json:"src_ip"`
	DstIP       string    `json:"dst_ip"`
	SrcPort     int       `json:"src_port"`
	DstPort     int       `json:"dst_port"`
	Protocol    string    `json:"protocol"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
	PacketCount int64     `json:"packet_count"`
	ByteCount   int64     `json:"byte_count"`
	SrcMAC      string    `json:"src_mac,omitempty"`
	DstMAC      string    `json:"dst_mac,omitempty"`
	TLSVersion  string    `json:"tls_version,omitempty"`
	TLSSNI      string    `json:"tls_sni,omitempty"`
	DNSQueries  []string  `json:"dns_queries,omitempty"`
	AppProtocol string    `json:"app_protocol,omitempty"`
	IATAvgMs    float64   `json:"iat_avg_ms"`
	IATStdDevMs float64   `json:"iat_std_dev_ms"`
	Score       float64   `json:"score"`
	CreatedAt   time.Time `json:"created_at"`
}

type FlowDetail struct {
	FlowResponse
	PayloadSample []byte             `json:"-"`
	PacketSamples []PacketSampleItem `json:"packet_samples,omitempty"`
}

type PacketSampleItem struct {
	Timestamp time.Time `json:"ts"`
	Size      int       `json:"size"`
}
