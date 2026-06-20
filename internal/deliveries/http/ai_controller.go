package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Dokito555/mizuki/internal/constants"
	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/Dokito555/mizuki/internal/services/ai"
	"github.com/Dokito555/mizuki/internal/services/flow"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AIController struct {
	aiEngine         *ai.AIEngine
	flowService      *flow.FlowService
	flowRepo         repositories.FlowRepository
	flowAIRepo       repositories.FlowAIRepository
	batchRepo        repositories.UploadAIBatchRepository
	log              *logrus.Logger
}

func NewAIController(
	aiEngine *ai.AIEngine,
	flowService *flow.FlowService,
	flowRepo repositories.FlowRepository,
	flowAIRepo repositories.FlowAIRepository,
	batchRepo repositories.UploadAIBatchRepository,
	log *logrus.Logger,
) *AIController {
	return &AIController{
		aiEngine:    aiEngine,
		flowService: flowService,
		flowRepo:    flowRepo,
		flowAIRepo:  flowAIRepo,
		batchRepo:   batchRepo,
		log:         log,
	}
}

func (c *AIController) GetAIAnalysis(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		BadRequest(ctx, "invalid flow id")
		return
	}

	flow, err := c.flowRepo.FindByID(ctx.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, constants.ErrFlowNotFound) {
			NotFound(ctx, "flow not found")
			return
		}
		c.log.Errorf("get flow for AI analysis: %v", err)
		InternalError(ctx, "failed to get flow")
		return
	}

	flowAI, err := c.flowAIRepo.FindByFlowID(ctx.Request.Context(), flow.ID)
	if err != nil {
		c.log.Errorf("get AI analysis: %v", err)
		InternalError(ctx, "failed to get AI analysis")
		return
	}

	if flowAI == nil || flowAI.Analysis == nil {
		WriteJSON(ctx, http.StatusOK, gin.H{
			"flow_id": id,
			"status":  "not_analyzed",
		})
		return
	}

	var result ai.AnalysisResult
	if err := json.Unmarshal(*flowAI.Analysis, &result); err != nil {
		c.log.Errorf("unmarshal AI analysis: %v", err)
		InternalError(ctx, "failed to parse AI analysis")
		return
	}

	WriteJSON(ctx, http.StatusOK, gin.H{
		"flow_id":      id,
		"status":       "analyzed",
		"model":        flow.AIModel,
		"analyzed_at":  flow.AIAnalyzedAt,
		"analysis":     result,
	})
}

func (c *AIController) AnalyzeFlow(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		BadRequest(ctx, "invalid flow id")
		return
	}

	if !c.aiEngine.IsEnabled() {
		WriteError(ctx, http.StatusServiceUnavailable, "AI analysis is not enabled")
		return
	}

	flow, err := c.flowRepo.FindByID(ctx.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, constants.ErrFlowNotFound) {
			NotFound(ctx, "flow not found")
			return
		}
		c.log.Errorf("find flow for AI analyze: %v", err)
		InternalError(ctx, "failed to get flow")
		return
	}

	result, err := c.aiEngine.AnalyzeFlowSync(ctx.Request.Context(), flow, flow.RawFileID)
	if err != nil {
		c.log.Errorf("AI analysis failed: %v", err)
		InternalError(ctx, "AI analysis failed")
		return
	}

	WriteJSON(ctx, http.StatusOK, gin.H{
		"flow_id":     id,
		"model":       result.ModelUsed,
		"is_fallback": result.IsFallback,
		"analysis":    result,
	})
}

func (c *AIController) AnalyzeUploadBatch(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		BadRequest(ctx, "invalid upload id")
		return
	}

	if !c.aiEngine.IsEnabled() {
		WriteError(ctx, http.StatusServiceUnavailable, "AI analysis is not enabled")
		return
	}

	existing, err := c.batchRepo.FindByUploadID(ctx.Request.Context(), uint(id))
	if err != nil {
		c.log.Errorf("find batch: %v", err)
		InternalError(ctx, "failed to check batch status")
		return
	}
	if existing != nil && existing.Status == constants.BatchRunning {
		WriteError(ctx, http.StatusConflict, "batch AI analysis already in progress")
		return
	}

	now := time.Now()
	batch := &entities.UploadAIBatch{
		UploadID:  uint(id),
		Status:    constants.BatchRunning,
		StartedAt: &now,
	}
	if existing != nil {
		batch.ID = existing.ID
		batch.CreatedAt = existing.CreatedAt
	}
	if err := c.batchRepo.CreateOrUpdate(ctx.Request.Context(), batch); err != nil {
		c.log.Errorf("create batch: %v", err)
		InternalError(ctx, "failed to start batch")
		return
	}

	go func() {
		bCtx := context.Background()
		flows, total, err := c.flowRepo.FindAll(bCtx, models.FlowFilter{
			UploadID: uint(id),
			PageSize: 10000,
		})
		if err != nil {
			c.log.Errorf("fetch flows for batch: %v", err)
			batch.Status = constants.BatchError
			c.batchRepo.CreateOrUpdate(bCtx, batch)
			return
		}

		batch.TotalCount = int(total)
		c.batchRepo.CreateOrUpdate(bCtx, batch)

		for _, f := range flows {
			c.aiEngine.EnrichFlow(bCtx, f.ID, uint(id))
			batch.ProcessedCount++
		}
		batch.Status = constants.BatchDone
		now := time.Now()
		batch.CompletedAt = &now
		c.batchRepo.CreateOrUpdate(bCtx, batch)
	}()

	WriteJSON(ctx, http.StatusAccepted, gin.H{
		"message":   "batch AI analysis queued",
		"upload_id": id,
		"batch_id":  batch.ID,
	})
}
