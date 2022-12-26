package cmd

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"shorturl/internal/auth"
	"shorturl/internal/handler"
	"shorturl/internal/params"
	"shorturl/internal/shorturl"
	"time"
)

type ServeCmd struct {
	params params.Params
	logger *zap.Logger
}

func NewServeCmd(params params.Params, logger *zap.Logger) *ServeCmd {
	return &ServeCmd{
		params: params,
		logger: logger,
	}
}

func (sc ServeCmd) Exec(_ *cobra.Command, _ []string) {
	engine := gin.Default()
	shortUrlRepo := shorturl.NewRepo(sc.params)
	contextAuth := auth.NewBearerSharedTokenContextAuth(sc.params)

	engine.Use(ginzap.Ginzap(sc.logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(sc.logger, true))

	handler.HandleIndex(engine)
	handler.HandleAddUrl(engine, shortUrlRepo, contextAuth)
	handler.HandleViewUrl(engine, shortUrlRepo)
	handler.HandleRedirect(engine, shortUrlRepo)

	listenAddr := sc.params.GetWithDefault(params.ListenAddr, ":8080")
	if err := engine.Run(listenAddr); err != nil {
		sc.logger.Error("Server start failed", zap.Error(err))
	}
}
