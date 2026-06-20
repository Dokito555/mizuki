package ai

import (
	"context"

	"github.com/Dokito555/mizuki/internal/entities"
)

type Provider interface {
	Analyze(ctx context.Context, flow *entities.Flow, relatedFlows []entities.Flow) (*AnalysisResult, error)
	Name() string
}

type AnalysisResult struct {
	ModelUsed    string            `json:"model_used"`
	Confidence   float64           `json:"confidence"`
	Narrative    string            `json:"narrative"`
	MitreIDs     []string          `json:"mitre_ids"`
	Attribution  string            `json:"attribution"`
	Remediation  []string          `json:"remediation"`
	IsFallback   bool              `json:"is_fallback"`
	Correlations []FlowCorrelation `json:"correlations,omitempty"`
}

type FlowCorrelation struct {
	FlowIDs     []uint `json:"flow_ids"`
	Pattern     string `json:"pattern"`
	Description string `json:"description"`
}

type AnalysisRequest struct {
	Flow     *entities.Flow
	Related  []entities.Flow
	UploadID uint
}
