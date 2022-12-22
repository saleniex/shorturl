package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"shorturl/cmd"
)

func main() {
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
		Run:   cmd.ServeCmd,
	}
	rootCmd.AddCommand(serveCmd)

	consumeStatsCmd := &cobra.Command{
		Use:   "consume-stats",
		Short: "AMQP stats consumer",
		Long:  "Consume URL access stats sent via AMQP",
		Run:   cmd.ConsumeStatsCmd,
	}
	rootCmd.AddCommand(consumeStatsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error %s", err)
	}
}
