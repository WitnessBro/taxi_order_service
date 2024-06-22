package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"taxi_order_service/cmd/runserver"
)

var RootCmd = &cobra.Command{
	Use:   "taxi_order_service",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		slog.Error("can't run command", "error", err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(runserver.Cmd)
}
