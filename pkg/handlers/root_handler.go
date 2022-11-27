package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleRoot handles requests for GET /
func HandleRoot(engine *gin.Engine) {
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"app": "url-shortener",
		})
	})
}
