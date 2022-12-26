package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"shorturl/cmd"
	"shorturl/internal/params"
)

func main() {
	param := params.NewEnvParams()
	logger, logErr := zap.NewProduction()
	if logErr != nil {
		panic(fmt.Sprintf("cannot fire-up logger: %s", logErr))
	}
	rootCmd := &cobra.Command{
		Use:   "shorturl",
		Short: "Short URL service",
		Long:  "Service for serving shortened URLs. Supports various URL repositories and performance optimizations",
		Run:   cmd.RootCmd,
	}

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve web service",
		Long:  "Application wbe service and API",
		Run:   cmd.NewServeCmd(param, logger).Exec,
	}
	rootCmd.AddCommand(serveCmd)

	consumeStatsCmd := &cobra.Command{
		Use:   "consume-stats",
		Short: "AMQP stats consumer",
		Long:  "Consume URL access stats sent via AMQP",
		Run:   cmd.NewConsumeStatsCmd(param, logger).Exec,
	}
	rootCmd.AddCommand(consumeStatsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error %s", err)
	}
}
