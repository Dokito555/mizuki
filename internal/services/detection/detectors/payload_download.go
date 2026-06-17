package detectors

import (
	"context"

	"github.com/Dokito555/mizuki/internal/entities"
)

type PayloadDownloadDetector struct{}

func NewPayloadDownloadDetector() *PayloadDownloadDetector {
	return &PayloadDownloadDetector{}
}

func (d *PayloadDownloadDetector) Name() string {
	return "payload_download"
}

func (d *PayloadDownloadDetector) PreProcess(ctx context.Context, allFlows []entities.Flow) {}

func (d *PayloadDownloadDetector) DetectFlow(ctx context.Context, flow *entities.Flow) (float64, string) {
	if flow.ByteCount < 500_000 {
		return 0, ""
	}

	avgPacketSize := float64(flow.ByteCount) / float64(flow.PacketCount)
	isBulkTransfer := avgPacketSize > 1000

	switch {
	case flow.ByteCount >= 50_000_000:
		return 85, "Large payload download (>50MB)"
	case flow.ByteCount >= 10_000_000:
		if isBulkTransfer {
			return 65, "Suspicious download (10-50MB)"
		}
		return 50, "Large data transfer (10-50MB)"
	case flow.ByteCount >= 1_000_000:
		if isBulkTransfer {
			return 40, "Bulk data transfer"
		}
		return 20, ""
	}

	return 0, ""
}
