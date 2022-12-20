package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/shorturl"
)

// HandleViewUrl handles request for GET /view/:id
func HandleViewUrl(e *gin.Engine, repo shorturl.Repo) {
	e.GET("/view/:id", func(context *gin.Context) {
		var uri shorturl.ShortIdUri
		if err := context.ShouldBindUri(&uri); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"errorMessage": err})
			return
		}

		url := repo.Find(uri.ShortId)
		if url == "" {
			context.Status(http.StatusNotFound)
			return
		}

		stats, statsErr := repo.ShortUrlAccessStats(uri.ShortId)
		if statsErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"errorMessage": statsErr.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"id":          uri.ShortId,
			"shortIdUri":  repo.Find(uri.ShortId),
			"accessCount": stats.Count,
		})
	})
}
