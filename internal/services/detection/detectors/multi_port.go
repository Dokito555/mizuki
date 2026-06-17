package detectors

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/entities"
)

type MultiPortDetector struct {
	portCounts map[string]int
}

func NewMultiPortDetector() *MultiPortDetector {
	return &MultiPortDetector{}
}

func (d *MultiPortDetector) Name() string {
	return "multi_port_c2"
}

func (d *MultiPortDetector) PreProcess(ctx context.Context, allFlows []entities.Flow) {
	d.portCounts = computePortCounts(allFlows)
}

func (d *MultiPortDetector) DetectFlow(ctx context.Context, flow *entities.Flow) (float64, string) {
	key := fmt.Sprintf("%s|%s", flow.SrcIP, flow.DstIP)
	ports := d.portCounts[key]

	switch {
	case ports >= 15:
		return 95, "Multi-port C2 — aggressive port scanning"
	case ports >= 8:
		return 80, "Multi-port C2 — extensive port usage"
	case ports >= 5:
		return 65, "Multi-port C2 — multiple ports"
	case ports >= 3:
		return 40, "Unusual port diversity"
	}

	return 0, ""
}

func computePortCounts(allFlows []entities.Flow) map[string]int {
	counts := make(map[string]int)
	seen := make(map[string]map[int]bool)

	for _, f := range allFlows {
		key := fmt.Sprintf("%s|%s", f.SrcIP, f.DstIP)
		if seen[key] == nil {
			seen[key] = make(map[int]bool)
		}
		seen[key][f.DstPort] = true
	}

	for key, ports := range seen {
		counts[key] = len(ports)
	}

	return counts
}
