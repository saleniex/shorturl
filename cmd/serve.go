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
	repo   shorturl.Repo
}

func NewServeCmd(params params.Params, logger *zap.Logger, repo shorturl.Repo) *ServeCmd {
	return &ServeCmd{
		params: params,
		logger: logger,
		repo:   repo,
	}
}

func (sc ServeCmd) Exec(_ *cobra.Command, _ []string) {
	engine := gin.Default()
	contextAuth := auth.NewBearerSharedTokenContextAuth(sc.params)

	engine.Use(ginzap.Ginzap(sc.logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(sc.logger, true))

	handler.HandleIndex(engine)
	handler.HandleAddUrl(engine, sc.repo, contextAuth)
	handler.HandleViewUrl(engine, sc.repo)
	handler.HandleRedirect(engine, sc.repo)

	listenAddr := sc.params.GetWithDefault(params.ListenAddr, ":8080")
	if err := engine.Run(listenAddr); err != nil {
		sc.logger.Error("Server start failed", zap.Error(err))
	}
}
