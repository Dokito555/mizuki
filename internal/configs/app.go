package configs

import (
	"github.com/Dokito555/mizuki/internal/deliveries/http"
	"github.com/Dokito555/mizuki/internal/deliveries/http/route"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/Dokito555/mizuki/internal/services/ai"
	"github.com/Dokito555/mizuki/internal/services/detection"
	"github.com/Dokito555/mizuki/internal/services/flow"
	"github.com/Dokito555/mizuki/internal/services/pcap"
	"github.com/Dokito555/mizuki/internal/services/upload"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App    *gin.Engine
	DB     *gorm.DB
	Log    *logrus.Logger
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	maxFileSize := int64(config.Config.GetInt("MAX_FILE_SIZE_MB")) * 1024 * 1024
	if maxFileSize <= 0 {
		maxFileSize = 500 * 1024 * 1024
	}

	pcapEngine := pcap.NewEngine()

	flowRepo := repositories.NewFlowRepository(config.DB)
	flowAIRepo := repositories.NewFlowAIRepository(config.DB)
	uploadRepo := repositories.NewUploadRepository(config.DB)
	batchRepo := repositories.NewUploadAIBatchRepository(config.DB)

	aiEngine := ai.NewAIEngine(flowRepo, flowAIRepo, uploadRepo, config.Config, config.Log)

	flowService := flow.NewFlowService(flowRepo, uploadRepo, config.Log)
	detectionEngine := detection.NewDetectionEngine(flowRepo, aiEngine, config.Log)
	uploadService := upload.NewUploadService(uploadRepo, flowRepo, pcapEngine, detectionEngine, config.Config, config.Log, maxFileSize)

	healthController := http.NewHealthController(config.Log)
	pcapController := http.NewPcapController(uploadService, flowService, config.Log, maxFileSize)
	aiController := http.NewAIController(aiEngine, flowService, flowRepo, flowAIRepo, batchRepo, config.Log)

	routeConfig := route.RouteConfig{
		App:              config.App,
		HealthController: healthController,
		PcapController:   pcapController,
		AIController:     aiController,
		Log:              config.Log,
	}

	routeConfig.Setup()
}
