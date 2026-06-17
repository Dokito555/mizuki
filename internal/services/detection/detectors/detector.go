package detectors

import (
	"context"

	"github.com/Dokito555/mizuki/internal/entities"
)

type Detector interface {
	Name() string
	PreProcess(ctx context.Context, allFlows []entities.Flow)
	DetectFlow(ctx context.Context, flow *entities.Flow) (score float64, threatLabel string)
}
