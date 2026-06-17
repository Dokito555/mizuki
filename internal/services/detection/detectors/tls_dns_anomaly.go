package detectors

import (
	"context"
	"math"

	"github.com/Dokito555/mizuki/internal/entities"
)

type TLSDNSAnomalyDetector struct{}

func NewTLSDNSAnomalyDetector() *TLSDNSAnomalyDetector {
	return &TLSDNSAnomalyDetector{}
}

func (d *TLSDNSAnomalyDetector) Name() string {
	return "tls_dns_anomaly"
}

func (d *TLSDNSAnomalyDetector) PreProcess(ctx context.Context, allFlows []entities.Flow) {}

func (d *TLSDNSAnomalyDetector) DetectFlow(ctx context.Context, flow *entities.Flow) (float64, string) {
	var score float64
	var threats []string

	if flow.TLSVersion != "" && flow.TLSSNI == "" {
		score += 70
		threats = append(threats, "TLS without SNI")
	}

	if flow.AppProtocol == "DNS" && flow.PacketCount > 500 {
		score += 80
		threats = append(threats, "DNS flood / high volume")
	} else if flow.AppProtocol == "DNS" && flow.PacketCount > 100 {
		score += 30
		threats = append(threats, "Elevated DNS activity")
	}

	if score <= 0 {
		return 0, ""
	}

	label := ""
	for i, t := range threats {
		if i > 0 {
			label += "; "
		}
		label += t
	}

	return math.Min(score, 100), label
}
