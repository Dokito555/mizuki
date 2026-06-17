package route

import (
	"github.com/Dokito555/mizuki/internal/deliveries/http"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App               *gin.Engine
	HealthController  *http.HealthController
}

func (c *RouteConfig) Setup() {
	c.SetupPublicRoute()
}

func (c *RouteConfig) SetupPublicRoute() {
	c.App.POST("/api/healthcheck", c.HealthController.HealthCheck)
}
