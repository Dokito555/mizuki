package ai

import (
	"encoding/json"
	"math"
	"regexp"
	"strings"
)

var jsonRe = regexp.MustCompile(`\{[^{}]*\}`)

type ollamaResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

type borderlineResult struct {
	IsMalicious bool    `json:"is_malicious"`
	Confidence  float64 `json:"confidence"`
	Reasoning   string  `json:"reasoning"`
}

type narrativeResult struct {
	Narrative   string   `json:"narrative"`
	MitreIDs    []string `json:"mitre_ids"`
	Attribution string   `json:"attribution"`
}

type correlationResult struct {
	Patterns []struct {
		Description string  `json:"description"`
		Flows       []uint  `json:"flows"`
		Confidence  float64 `json:"confidence"`
	} `json:"patterns"`
}

type remediationResult struct {
	Remediation []string `json:"remediation"`
	Priority    string   `json:"priority"`
}

func ParseOllamaResponse(body []byte) (string, error) {
	var resp ollamaResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}
	return resp.Message.Content, nil
}

func extractJSON(raw string) string {
	m := jsonRe.FindString(raw)
	if m == "" {
		match := strings.Index(raw, "{")
		end := strings.LastIndex(raw, "}")
		if match >= 0 && end > match {
			return raw[match : end+1]
		}
		return ""
	}
	return m
}

func ClampConfidence(v float64) float64 {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return 0
	}
	if v < 0 {
		return 0
	}
	if v > 1.0 {
		return 1.0
	}
	return v
}

func ParseBorderline(raw string) (*AnalysisResult, error) {
	jsonStr := extractJSON(raw)
	if jsonStr == "" {
		return nil, nil
	}

	var br borderlineResult
	if err := json.Unmarshal([]byte(jsonStr), &br); err != nil {
		return nil, err
	}

	return &AnalysisResult{
		Confidence: ClampConfidence(br.Confidence),
		Narrative:  br.Reasoning,
	}, nil
}

func ParseNarrative(raw string, mitreValidator func([]string) []string) (*AnalysisResult, error) {
	jsonStr := extractJSON(raw)
	if jsonStr == "" {
		return nil, nil
	}

	var nr narrativeResult
	if err := json.Unmarshal([]byte(jsonStr), &nr); err != nil {
		return nil, err
	}

	if mitreValidator != nil {
		nr.MitreIDs = mitreValidator(nr.MitreIDs)
	}

	return &AnalysisResult{
		Narrative:   nr.Narrative,
		MitreIDs:    nr.MitreIDs,
		Attribution: nr.Attribution,
	}, nil
}

func ParseCorrelation(raw string) ([]FlowCorrelation, error) {
	jsonStr := extractJSON(raw)
	if jsonStr == "" {
		return nil, nil
	}

	var cr correlationResult
	if err := json.Unmarshal([]byte(jsonStr), &cr); err != nil {
		return nil, err
	}

	var correlations []FlowCorrelation
	for _, p := range cr.Patterns {
		correlations = append(correlations, FlowCorrelation{
			FlowIDs:     p.Flows,
			Pattern:     p.Description,
			Description: p.Description,
		})
	}
	return correlations, nil
}

func ParseRemediation(raw string) ([]string, error) {
	jsonStr := extractJSON(raw)
	if jsonStr == "" {
		return nil, nil
	}

	var rr remediationResult
	if err := json.Unmarshal([]byte(jsonStr), &rr); err != nil {
		return nil, err
	}

	return rr.Remediation, nil
}
