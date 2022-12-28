package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/shorturl"
)

type ViewUlr struct {
	repo shorturl.Repo
}

func NewViewUrl(repo shorturl.Repo) *ViewUlr {
	return &ViewUlr{
		repo: repo,
	}
}

func (u ViewUlr) Handle(context *gin.Context) {
	var uri shorturl.ShortIdUri
	if err := context.ShouldBindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errorMessage": err})
		return
	}

	url := u.repo.Find(uri.ShortId)
	if url == "" {
		context.Status(http.StatusNotFound)
		return
	}

	stats, statsErr := u.repo.ShortUrlAccessStats(uri.ShortId)
	if statsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errorMessage": statsErr.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":          uri.ShortId,
		"shortIdUri":  u.repo.Find(uri.ShortId),
		"accessCount": stats.Count,
	})
}
