package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {
}

func (i Index) Handle(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"app": "url-shortener",
	})
}
