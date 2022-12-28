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

	engine.GET("/", handler.Index{}.Handle)

	redirectHandler := handler.NewRedirect(sc.repo)
	engine.GET("/go/:id", redirectHandler.Handle)

	viewHandler := handler.NewViewUrl(sc.repo)
	engine.GET("/view/:id", viewHandler.Handle)

	addUrlHandler := handler.NewAddUrl(sc.repo, contextAuth, shorturl.NewSimpleIdGenerator())
	engine.POST("/", addUrlHandler.Handle)

	listenAddr := sc.params.GetWithDefault(params.ListenAddr, ":8080")
	if err := engine.Run(listenAddr); err != nil {
		sc.logger.Error("Server start failed", zap.Error(err))
	}
}
