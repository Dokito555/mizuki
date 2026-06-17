package detectors

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/entities"
)

var portExpectations = map[string][]int{
	"HTTP":   {80, 8080, 8000},
	"HTTPS":  {443},
	"SMTP":   {25, 465, 587},
	"SSH":    {22},
	"DNS":    {53},
	"FTP":    {21},
	"TELNET": {23},
}

type ProtocolAnomalyDetector struct{}

func NewProtocolAnomalyDetector() *ProtocolAnomalyDetector {
	return &ProtocolAnomalyDetector{}
}

func (d *ProtocolAnomalyDetector) Name() string {
	return "protocol_anomaly"
}

func (d *ProtocolAnomalyDetector) PreProcess(ctx context.Context, allFlows []entities.Flow) {}

func (d *ProtocolAnomalyDetector) DetectFlow(ctx context.Context, flow *entities.Flow) (float64, string) {
	if flow.AppProtocol == "" {
		return 0, ""
	}

	expected, ok := portExpectations[flow.AppProtocol]
	if !ok {
		return 0, ""
	}

	for _, p := range expected {
		if flow.DstPort == p {
			return 0, ""
		}
	}

	if flow.TLSVersion != "" && flow.AppProtocol != "HTTPS" {
		return 90, "Encrypted traffic on unexpected port"
	}

	return 70, "Protocol " + flow.AppProtocol + " on non-standard port " + fmt.Sprintf("%d", flow.DstPort)
}
