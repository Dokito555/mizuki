package detectors

import (
	"context"

	"github.com/Dokito555/mizuki/internal/entities"
)

type BeaconingDetector struct{}

func NewBeaconingDetector() *BeaconingDetector {
	return &BeaconingDetector{}
}

func (d *BeaconingDetector) Name() string {
	return "c2_beaconing"
}

func (d *BeaconingDetector) PreProcess(ctx context.Context, allFlows []entities.Flow) {}

func (d *BeaconingDetector) DetectFlow(ctx context.Context, flow *entities.Flow) (float64, string) {
	if flow.PacketCount < 5 {
		return 0, ""
	}
	if flow.IATAvgMs <= 0 {
		return 0, ""
	}

	cv := flow.IATStdDevMs / flow.IATAvgMs
	isSmallPacket := flow.ByteCount/flow.PacketCount < 200
	isOneSided := flow.ByteCount < 500

	switch {
	case cv < 0.3:
		return 90, "C2 Beaconing — highly periodic"
	case cv < 0.5:
		if isSmallPacket && isOneSided {
			return 85, "C2 Beaconing — periodic with small payload"
		}
		return 60, "C2 Beaconing — moderately periodic"
	case cv < 0.8:
		if isSmallPacket {
			return 40, "Suspicious periodicity"
		}
		return 20, ""
	}

	return 0, ""
}
