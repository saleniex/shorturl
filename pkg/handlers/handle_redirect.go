package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/pkg/repo"
)

func HandleRedirect(e *gin.Engine, repo repo.ShortUrlRepo) {
	e.GET("/go/:id", func(context *gin.Context) {
		var uri ShortIdUri
		if err := context.ShouldBindUri(&uri); err != nil {
			context.Status(http.StatusBadRequest)
			return
		}
		url := repo.Find(uri.ShortId)
		if url == "" {
			context.Status(http.StatusNotFound)
			return
		}

		_ = repo.LogAccess(uri.ShortId, context.ClientIP())

		context.Redirect(http.StatusFound, url)
	})
}