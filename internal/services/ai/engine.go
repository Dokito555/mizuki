package ai

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AIEngine struct {
	provider   Provider
	fallback   *FallbackProvider
	queue      *AIEnrichmentQueue
	flowRepo   repositories.FlowRepository
	flowAIRepo repositories.FlowAIRepository
	config     *viper.Viper
	log        *logrus.Logger
	enabled    bool
}

func NewAIEngine(
	flowRepo repositories.FlowRepository,
	flowAIRepo repositories.FlowAIRepository,
	uploadRepo repositories.UploadRepository,
	config *viper.Viper,
	log *logrus.Logger,
) *AIEngine {
	enabled := config.GetBool("AI_ENABLED")

	var provider Provider
	var queue *AIEnrichmentQueue

	if enabled {
		ollamaCfg := OllamaConfig{
			URL:     config.GetString("OLLAMA_URL"),
			Model:   config.GetString("OLLAMA_MODEL"),
			Timeout: time.Duration(config.GetInt("AI_TIMEOUT_SECONDS")) * time.Second,
			NumCtx:  config.GetInt("OLLAMA_NUM_CTX"),
		}
		if ollamaCfg.URL == "" {
			ollamaCfg.URL = "http://localhost:11434"
		}
		if ollamaCfg.Model == "" {
			ollamaCfg.Model = "qwen2.5:3b"
		}
		if ollamaCfg.Timeout <= 0 {
			ollamaCfg.Timeout = 60 * time.Second
		}
		if ollamaCfg.NumCtx <= 0 {
			ollamaCfg.NumCtx = 4096
		}

		provider = NewOllamaProvider(ollamaCfg)
		concurrency := config.GetInt("AI_CONCURRENCY")
		if concurrency < 1 {
			concurrency = 3
		}
		queue = NewAIEnrichmentQueue(concurrency, flowRepo, flowAIRepo, uploadRepo, []Provider{provider}, log)
	}

	return &AIEngine{
		provider:   provider,
		fallback:   NewFallbackProvider(),
		queue:      queue,
		flowRepo:   flowRepo,
		flowAIRepo: flowAIRepo,
		config:     config,
		log:        log,
		enabled:    enabled,
	}
}

func (e *AIEngine) IsEnabled() bool {
	return e.enabled
}

func (e *AIEngine) EnrichFlow(ctx context.Context, flowID uint, uploadID uint) {
	if !e.enabled || e.queue == nil {
		return
	}
	e.queue.Enqueue(EnrichmentJob{
		FlowID:   flowID,
		UploadID: uploadID,
	})
}

func (e *AIEngine) AnalyzeFlowSync(ctx context.Context, flow *entities.Flow, uploadID uint) (*AnalysisResult, error) {
	if !e.enabled || e.provider == nil {
		return e.fallback.Analyze(ctx, flow, nil)
	}

	related, _, _ := e.flowRepo.FindAll(ctx, models.FlowFilter{
		UploadID: uploadID,
		SrcIP:    flow.SrcIP,
	})

	result, err := e.provider.Analyze(ctx, flow, related)
	if err != nil {
		e.log.WithError(err).Warn("AI provider failed, using fallback")
		return e.fallback.Analyze(ctx, flow, related)
	}

	resultJSON, _ := json.Marshal(result)
	if err := e.flowAIRepo.Save(ctx, flow.ID, resultJSON); err != nil {
		e.log.WithError(err).Warn("failed to persist AI analysis")
	}
	if err := e.flowRepo.MarkAIAnalyzed(ctx, flow.ID, result.ModelUsed); err != nil {
		e.log.WithError(err).Warn("failed to mark AI analyzed")
	}

	return result, nil
}

func (e *AIEngine) Shutdown() {
	if e.queue != nil {
		e.queue.Stop()
	}
}
