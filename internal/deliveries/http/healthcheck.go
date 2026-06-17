package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	log *logrus.Logger
}

func NewHealthController(log *logrus.Logger) *HealthController {
	return &HealthController{
		log: log,
	}
}

func (hc *HealthController) HealthCheck(c *gin.Context) {
	hc.log.Info("Health check")
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
