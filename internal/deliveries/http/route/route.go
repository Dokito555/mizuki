package route

import (
	"github.com/Dokito555/mizuki/internal/deliveries/http"
	"github.com/Dokito555/mizuki/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App              *gin.Engine
	HealthController *http.HealthController
	PcapController   *http.PcapController
}

func (c *RouteConfig) Setup() {
	c.App.Use(middlewares.CORS())
	c.App.Use(middlewares.Recovery(nil))

	api := c.App.Group("/api")
	{
		api.POST("/healthcheck", c.HealthController.HealthCheck)

		pcap := api.Group("/pcap")
		{
			pcap.POST("/upload", c.PcapController.Upload)
		}

		uploads := api.Group("/uploads")
		{
			uploads.GET("", c.PcapController.ListUploads)
			uploads.GET("/:id", c.PcapController.GetUpload)
			uploads.POST("/:id/reparse", c.PcapController.Reparse)
		}

		flows := api.Group("/flows")
		{
			flows.GET("", c.PcapController.ListFlows)
			flows.GET("/:id", c.PcapController.GetFlow)
		}
	}
}
