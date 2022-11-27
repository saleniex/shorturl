package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/pkg/app"
	"shorturl/pkg/auth"
	"shorturl/pkg/repo"
)

// HandleAddUrl handles request for POST /
//
// Content type is expected to be "application/json" while body contains JSON object with two parameters:
// "url" and "shortId"
func HandleAddUrl(engine *gin.Engine, urlRepo repo.ShortUrlRepo, auth auth.ContextAuth) {
	engine.POST("/", func(context *gin.Context) {
		if err := auth.Authorize(context); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
		var shortUrl app.ShortUrl
		if err := context.ShouldBindJSON(&shortUrl); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errorMessage": "cannot bind input params " + err.Error(),
			})
			return
		}

		if err := urlRepo.StoreUrl(shortUrl); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errorMessage": "cannot store url " + err.Error(),
			})
			return
		}

		context.Status(http.StatusCreated)
	})
}
