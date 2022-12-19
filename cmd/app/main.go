package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shorturl/internal/auth"
	"shorturl/internal/handler"
	"shorturl/internal/params"
	"shorturl/internal/shorturl"
	"time"
)

func main() {
	parameters := params.NewEnvParams()
	engine := gin.Default()
	shortUrlRepo := shorturl.NewRepo(parameters)
	contextAuth := auth.NewBearerSharedTokenContextAuth(parameters)

	logger, _ := zap.NewProduction()
	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))

	handler.HandleIndex(engine)
	handler.HandleAddUrl(engine, shortUrlRepo, contextAuth)
	handler.HandleViewUrl(engine, shortUrlRepo)
	handler.HandleRedirect(engine, shortUrlRepo)

	listenAddr := parameters.GetWithDefault(params.ListenAddr, ":8080")
	if err := engine.Run(listenAddr); err != nil {
		logger.Error("Server start failed", zap.Error(err))
	}
}
