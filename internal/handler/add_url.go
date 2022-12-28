package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/auth"
	"shorturl/internal/shorturl"
)

type AddUrl struct {
	urlRepo shorturl.Repo
	auth    auth.ContextAuth
	idGen   shorturl.IdGenerator
}

func NewAddUrl(urlRepo shorturl.Repo, auth auth.ContextAuth, idGen shorturl.IdGenerator) *AddUrl {
	return &AddUrl{
		urlRepo: urlRepo,
		auth:    auth,
		idGen:   idGen,
	}
}

func (u AddUrl) Handle(context *gin.Context) {
	if err := u.auth.Authorize(context); err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}
	var shortUrl shorturl.ShortUrl
	if err := context.ShouldBindJSON(&shortUrl); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "cannot bind input params " + err.Error(),
		})
		return
	}

	if shortUrl.ShortId == "" {
		shortUrl.ShortId = u.idGen.Generate()
	}

	if err := u.urlRepo.StoreUrl(shortUrl); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "cannot store url " + err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         shortUrl.ShortId,
		"shortIdUri": shortUrl.Url,
	})
}
