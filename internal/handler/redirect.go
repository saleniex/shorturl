package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/shorturl"
)

type Redirect struct {
	repo shorturl.Repo
}

func NewRedirect(repo shorturl.Repo) *Redirect {
	return &Redirect{
		repo: repo,
	}
}

func (r Redirect) Handle(context *gin.Context) {
	var uri shorturl.ShortIdUri
	if err := context.ShouldBindUri(&uri); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	url := r.repo.Find(uri.ShortId)
	if url == "" {
		context.Status(http.StatusNotFound)
		return
	}

	_ = r.repo.LogAccess(uri.ShortId, context.ClientIP())

	context.Redirect(http.StatusFound, url)
}
