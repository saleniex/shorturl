package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RootCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("Use flag `-h` for more details how to use this application.\n")
}
