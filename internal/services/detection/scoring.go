package detection

import "math"

type DetectorResult struct {
	Name  string
	Score float64
	Label string
}

type ScoringEngine struct {
	weights map[string]float64
}

func NewScoringEngine() *ScoringEngine {
	return &ScoringEngine{
		weights: map[string]float64{
			"c2_beaconing":     0.35,
			"multi_port_c2":    0.25,
			"protocol_anomaly": 0.20,
			"payload_download": 0.15,
			"tls_dns_anomaly":  0.05,
		},
	}
}

func (s *ScoringEngine) FinalScore(results []DetectorResult) float64 {
	var total, weightSum float64
	for _, r := range results {
		if r.Score <= 0 {
			continue
		}
		w := s.weights[r.Name]
		if w <= 0 {
			continue
		}
		total += r.Score * w
		weightSum += w
	}
	if weightSum == 0 {
		return 0
	}
	return math.Round(total / weightSum)
}

func (s *ScoringEngine) CollectThreats(results []DetectorResult) []string {
	var threats []string
	for _, r := range results {
		if r.Score >= 50 && r.Label != "" {
			threats = append(threats, r.Label)
		}
	}
	return threats
}
