package configs

import (
	// "errors"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	return gin.New()
}

// func NewErrorHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Next()
// 		if len(ctx.Errors) > 0 {
// 			err := ctx.Errors.Last()
// 			code := http.StatusInternalServerError

// 			var ginErr *gin.Error
// 			if errors.As(err.Err, &ginErr) {
// 				code = int(ginErr.Type)
// 			}

// 			ctx.JSON(code, gin.H{
// 				"errors": err.Error(),
// 			})
// 		}
// 	}
// }