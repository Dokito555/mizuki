package detection

import (
	"context"
	"fmt"

	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/Dokito555/mizuki/internal/services/detection/detectors"
	"github.com/sirupsen/logrus"
)

const pageSize = 1000

type DetectionEngine struct {
	detectors []detectors.Detector
	flowRepo  repositories.FlowRepository
	scoring   *ScoringEngine
	log       *logrus.Logger
}

func NewDetectionEngine(
	flowRepo repositories.FlowRepository,
	log *logrus.Logger,
) *DetectionEngine {
	return &DetectionEngine{
		detectors: []detectors.Detector{
			detectors.NewBeaconingDetector(),
			detectors.NewMultiPortDetector(),
			detectors.NewProtocolAnomalyDetector(),
			detectors.NewPayloadDownloadDetector(),
			detectors.NewTLSDNSAnomalyDetector(),
		},
		flowRepo: flowRepo,
		scoring:  NewScoringEngine(),
		log:      log,
	}
}

func (e *DetectionEngine) AnalyzeUpload(ctx context.Context, uploadID uint) error {
	log := e.log.WithField("upload_id", uploadID)
	log.Info("starting analysis")

	page := 1
	var totalAnalyzed int
	var allFlows []entities.Flow

	for {
		flows, total, err := e.flowRepo.FindAll(ctx, models.FlowFilter{
			UploadID: uploadID,
			Page:     page,
			PageSize: pageSize,
		})
		if err != nil {
			return fmt.Errorf("detectionEngine.AnalyzeUpload fetch page %d: %w", page, err)
		}
		if len(flows) == 0 {
			break
		}
		allFlows = append(allFlows, flows...)
		totalAnalyzed += len(flows)
		if totalAnalyzed >= int(total) {
			break
		}
		page++
	}

	if len(allFlows) == 0 {
		log.Info("no flows to analyze")
		return nil
	}

	for _, d := range e.detectors {
		d.PreProcess(ctx, allFlows)
	}

	for i := range allFlows {
		e.analyzeFlow(ctx, &allFlows[i])
	}

	if err := e.flowRepo.UpdateScores(ctx, allFlows); err != nil {
		return fmt.Errorf("detectionEngine.AnalyzeUpload update: %w", err)
	}

	log.Infof("analysis complete: %d flows analyzed", totalAnalyzed)
	return nil
}

func (e *DetectionEngine) analyzeFlow(ctx context.Context, flow *entities.Flow) {
	var results []DetectorResult

	for _, d := range e.detectors {
		score, label := d.DetectFlow(ctx, flow)
		if score <= 0 {
			continue
		}
		results = append(results, DetectorResult{
			Name:  d.Name(),
			Score: score,
			Label: label,
		})
	}

	flow.Score = e.scoring.FinalScore(results)
	flow.Threats = e.scoring.CollectThreats(results)
}
