package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/shorturl"
)

func HandleRedirect(e *gin.Engine, repo shorturl.Repo) {
	e.GET("/go/:id", func(context *gin.Context) {
		var uri shorturl.ShortIdUri
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
