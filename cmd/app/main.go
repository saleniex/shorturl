package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shorturl/pkg/auth"
	"shorturl/pkg/handlers"
	"shorturl/pkg/params"
	"shorturl/pkg/repo"
	"time"
)

func main() {
	parameters := params.NewEnvParams()
	engine := gin.Default()
	shortUrlRepo := repo.NewRepo(parameters)
	contextAuth := auth.NewBearerSharedTokenContextAuth(parameters)

	logger, _ := zap.NewProduction()
	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))

	handlers.HandleRoot(engine)
	handlers.HandleAddUrl(engine, shortUrlRepo, contextAuth)
	handlers.HandleViewUrl(engine, shortUrlRepo)
	handlers.HandleRedirect(engine, shortUrlRepo)

	listenAddr := parameters.GetWithDefault(params.ListenAddr, ":8080")
	if err := engine.Run(listenAddr); err != nil {
		logger.Error("Server start failed", zap.Error(err))
	}
}
