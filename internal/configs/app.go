package configs

import (
	"github.com/Dokito555/mizuki/internal/deliveries/http"
	"github.com/Dokito555/mizuki/internal/deliveries/http/route"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	App      *gin.Engine
	Log      *logrus.Logger
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repo
	// setup services
	// setup controllers
	healthController := http.NewHealthController(config.Log)
	
	// setup middleware
	// route config
	routeConfig := route.RouteConfig{
		App:               config.App,
		HealthController:  healthController,
	}

	routeConfig.Setup()
}
