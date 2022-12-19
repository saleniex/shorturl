package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleIndex handles requests for GET /
func HandleIndex(engine *gin.Engine) {
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"app": "url-shortener",
		})
	})
}
