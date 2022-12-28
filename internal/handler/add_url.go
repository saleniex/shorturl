package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/auth"
	"shorturl/internal/shorturl"
)

// HandleAddUrl handles request for POST /
//
// Content type is expected to be "application/json" while body contains JSON object with two parameters:
// "url" and "shortId"
func HandleAddUrl(engine *gin.Engine, urlRepo shorturl.Repo, auth auth.ContextAuth, idGen shorturl.IdGenerator) {
	engine.POST("/", func(context *gin.Context) {
		if err := auth.Authorize(context); err != nil {
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
			shortUrl.ShortId = idGen.Generate()
		}

		if err := urlRepo.StoreUrl(shortUrl); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errorMessage": "cannot store url " + err.Error(),
			})
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"id":         shortUrl.ShortId,
			"shortIdUri": shortUrl.Url,
		})
	})
}
