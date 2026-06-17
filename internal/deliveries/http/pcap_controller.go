package http

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Dokito555/mizuki/internal/constants"
	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PcapController struct {
	uploadService *services.UploadService
	flowService   *services.FlowService
	log           *logrus.Logger
	maxFileSize   int64
}

func NewPcapController(
	uploadService *services.UploadService,
	flowService *services.FlowService,
	log *logrus.Logger,
	maxFileSize int64,
) *PcapController {
	return &PcapController{
		uploadService: uploadService,
		flowService:   flowService,
		log:           log,
		maxFileSize:   maxFileSize,
	}
}

func (c *PcapController) Upload(ctx *gin.Context) {
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, c.maxFileSize+1024)

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			WriteError(ctx, http.StatusRequestEntityTooLarge, "file exceeds maximum size")
			return
		}
		BadRequest(ctx, "file is required")
		return
	}
	defer file.Close()

	forceReparse := ctx.Query("force") == "true"

	result, err := c.uploadService.ProcessUpload(ctx.Request.Context(), file, header, forceReparse)
	if err != nil {
		if errors.Is(err, constants.ErrFileTooLarge) {
			WriteError(ctx, http.StatusRequestEntityTooLarge, "file exceeds maximum size")
			return
		}
		c.log.Errorf("upload failed: %v", err)
		InternalError(ctx, "upload processing failed")
		return
	}

	WriteJSON(ctx, http.StatusAccepted, result)
}

func (c *PcapController) GetUpload(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		BadRequest(ctx, "invalid upload id")
		return
	}

	result, err := c.uploadService.GetUploadByID(ctx.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, constants.ErrUploadNotFound) {
			NotFound(ctx, "upload not found")
			return
		}
		c.log.Errorf("get upload failed: %v", err)
		InternalError(ctx, "failed to get upload")
		return
	}

	WriteJSON(ctx, http.StatusOK, result)
}

func (c *PcapController) ListUploads(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	result, err := c.uploadService.ListUploads(ctx.Request.Context(), page, pageSize)
	if err != nil {
		c.log.Errorf("list uploads failed: %v", err)
		InternalError(ctx, "failed to list uploads")
		return
	}

	WritePaginated(ctx, http.StatusOK, result.Data, result.Meta)
}

func (c *PcapController) ListFlows(ctx *gin.Context) {
	var filter models.FlowFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		BadRequest(ctx, "invalid filter parameters")
		return
	}

	result, err := c.flowService.List(ctx.Request.Context(), filter)
	if err != nil {
		c.log.Errorf("list flows failed: %v", err)
		InternalError(ctx, "failed to list flows")
		return
	}

	WritePaginated(ctx, http.StatusOK, result.Data, result.Meta)
}

func (c *PcapController) AnalyzeUpload(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		BadRequest(ctx, "invalid upload id")
		return
	}

	go c.uploadService.AnalyzeUpload(context.Background(), uint(id))

	WriteJSON(ctx, http.StatusAccepted, gin.H{"message": "analysis started"})
}

func (c *PcapController) CancelUpload(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		BadRequest(ctx, "invalid upload id")
		return
	}

	if !c.uploadService.CancelUpload(uint(id)) {
		NotFound(ctx, "upload not found or already completed")
		return
	}
	WriteJSON(ctx, http.StatusOK, gin.H{"message": "upload cancelled"})
}

func (c *PcapController) GetFlow(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		BadRequest(ctx, "invalid flow id")
		return
	}

	result, err := c.flowService.GetByID(ctx.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, constants.ErrFlowNotFound) {
			NotFound(ctx, "flow not found")
			return
		}
		c.log.Errorf("get flow failed: %v", err)
		InternalError(ctx, "failed to get flow")
		return
	}

	WriteJSON(ctx, http.StatusOK, result)
}
