package configs

import (
	"github.com/Dokito555/mizuki/internal/deliveries/http"
	"github.com/Dokito555/mizuki/internal/deliveries/http/route"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/Dokito555/mizuki/internal/services"
	"github.com/Dokito555/mizuki/internal/services/detection"
	"github.com/Dokito555/mizuki/internal/services/pcap"
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
	uploadRepo := repositories.NewUploadRepository(config.DB)

	flowService := services.NewFlowService(flowRepo, uploadRepo, config.Log)
	detectionEngine := detection.NewDetectionEngine(flowRepo, config.Log)
	uploadService := services.NewUploadService(uploadRepo, flowRepo, pcapEngine, detectionEngine, config.Log, maxFileSize)

	healthController := http.NewHealthController(config.Log)
	pcapController := http.NewPcapController(uploadService, flowService, config.Log, maxFileSize)

	routeConfig := route.RouteConfig{
		App:              config.App,
		HealthController: healthController,
		PcapController:   pcapController,
		Log:              config.Log,
	}

	routeConfig.Setup()
}
