package ai

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/entities"
)

type FallbackProvider struct{}

func NewFallbackProvider() *FallbackProvider {
	return &FallbackProvider{}
}

func (p *FallbackProvider) Name() string {
	return "heuristic-fallback"
}

func (p *FallbackProvider) Analyze(ctx context.Context, flow *entities.Flow, _ []entities.Flow) (*AnalysisResult, error) {
	threats := []string(flow.Threats)
	narrative := "Heuristic detection"
	if len(threats) > 0 {
		narrative = fmt.Sprintf("Automated detection: %s (AI unavailable)", threats[0])
	}

	return &AnalysisResult{
		ModelUsed:   p.Name(),
		Confidence:  0.0,
		Narrative:   narrative,
		MitreIDs:    nil,
		Attribution: "",
		Remediation: nil,
		IsFallback:  true,
	}, nil
}
