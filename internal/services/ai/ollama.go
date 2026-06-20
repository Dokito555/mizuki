package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Dokito555/mizuki/internal/entities"
)

type OllamaConfig struct {
	URL         string
	Model       string
	Timeout     time.Duration
	NumCtx      int
}

type OllamaProvider struct {
	config      OllamaConfig
	httpClient  *http.Client
	mitreValidator func([]string) []string
}

func NewOllamaProvider(config OllamaConfig) *OllamaProvider {
	return &OllamaProvider{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		mitreValidator: SanitizeMitreIDs,
	}
}

func (p *OllamaProvider) Name() string {
	return p.config.Model
}

type ollamaChatRequest struct {
	Model    string        `json:"model"`
	Messages []ollamaMessage `json:"messages"`
	Stream   bool          `json:"stream"`
	Options  ollamaOptions `json:"options"`
}

type ollamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ollamaOptions struct {
	NumCtx int `json:"num_ctx"`
}

func (p *OllamaProvider) Analyze(ctx context.Context, flow *entities.Flow, related []entities.Flow) (*AnalysisResult, error) {
	result := &AnalysisResult{
		ModelUsed:  p.Name(),
		Confidence: 0,
		IsFallback: false,
	}

	systemBorder, userBorder := PromptTemplates["borderline"](flow, related)
	borderResp, err := p.callOllama(ctx, systemBorder, userBorder)
	if err != nil {
		return nil, fmt.Errorf("ollama borderline: %w", err)
	}

	borderResult, err := ParseBorderline(borderResp)
	if err != nil {
		return nil, fmt.Errorf("parse borderline: %w", err)
	}
	if borderResult == nil {
		return result, nil
	}

	result.Confidence = borderResult.Confidence
	result.Narrative = borderResult.Narrative

	systemNarr, userNarr := PromptTemplates["narrative"](flow, related)
	narrResp, err := p.callOllama(ctx, systemNarr, userNarr)
	if err != nil {
		return result, nil
	}

	narrResult, err := ParseNarrative(narrResp, p.mitreValidator)
	if err == nil && narrResult != nil {
		result.Narrative = narrResult.Narrative
		result.MitreIDs = narrResult.MitreIDs
		result.Attribution = narrResult.Attribution
	}

	systemRem, userRem := PromptTemplates["remediation"](flow, related)
	remResp, err := p.callOllama(ctx, systemRem, userRem)
	if err == nil {
		if remediations, err := ParseRemediation(remResp); err == nil {
			result.Remediation = remediations
		}
	}

	if len(related) > 0 {
		systemCorr, userCorr := PromptTemplates["correlation"](flow, related)
		corrResp, err := p.callOllama(ctx, systemCorr, userCorr)
		if err == nil {
			if correlations, err := ParseCorrelation(corrResp); err == nil {
				result.Correlations = correlations
			}
		}
	}

	return result, nil
}

func (p *OllamaProvider) callOllama(ctx context.Context, system, user string) (string, error) {
	reqBody := ollamaChatRequest{
		Model:  p.config.Model,
		Stream: false,
		Options: ollamaOptions{
			NumCtx: p.config.NumCtx,
		},
		Messages: []ollamaMessage{
			{Role: "system", Content: system},
			{Role: "user", Content: user},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.config.URL+"/api/chat", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama status %d: %s", resp.StatusCode, string(respBody))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	return ParseOllamaResponse(respBody)
}
