package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"data": data})
}

func WritePaginated(c *gin.Context, status int, data interface{}, meta interface{}) {
	c.JSON(status, gin.H{
		"data": data,
		"meta": meta,
	})
}

func WriteError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func WriteErrorWithCode(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func BadRequest(c *gin.Context, message string) {
	WriteError(c, http.StatusBadRequest, message)
}

func NotFound(c *gin.Context, message string) {
	WriteError(c, http.StatusNotFound, message)
}

func InternalError(c *gin.Context, message string) {
	WriteError(c, http.StatusInternalServerError, message)
}
