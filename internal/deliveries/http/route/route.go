package route

import (
	"github.com/Dokito555/mizuki/internal/deliveries/http"
	"github.com/Dokito555/mizuki/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RouteConfig struct {
	App              *gin.Engine
	HealthController *http.HealthController
	PcapController   *http.PcapController
	Log              *logrus.Logger
}

func (c *RouteConfig) Setup() {
	c.App.Use(middlewares.CORS())
	c.App.Use(middlewares.Recovery(c.Log))

	api := c.App.Group("/api")
	{
		api.GET("/healthcheck", c.HealthController.HealthCheck)

		pcap := api.Group("/pcap")
		{
			pcap.POST("/upload", c.PcapController.Upload)
		}

		uploads := api.Group("/uploads")
		{
			uploads.GET("", c.PcapController.ListUploads)
			uploads.GET("/:id", c.PcapController.GetUpload)
			uploads.POST("/:id/reparse", c.PcapController.Reparse)
			uploads.POST("/:id/cancel", c.PcapController.CancelUpload)
		}

		flows := api.Group("/flows")
		{
			flows.GET("", c.PcapController.ListFlows)
			flows.GET("/:id", c.PcapController.GetFlow)
		}
	}
}
