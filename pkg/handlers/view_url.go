package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/pkg/repo"
)

// HandleViewUrl handles request for GET /view/:id
func HandleViewUrl(e *gin.Engine, repo repo.ShortUrlRepo) {
	e.GET("/view/:id", func(context *gin.Context) {
		var uri ShortIdUri
		if err := context.ShouldBindUri(&uri); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"errorMessage": err})
			return
		}

		url := repo.Find(uri.ShortId)
		if url == "" {
			context.Status(http.StatusNotFound)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"id":         uri.ShortId,
			"shortIdUri": repo.Find(uri.ShortId),
		})
	})
}
