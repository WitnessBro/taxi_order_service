package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
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

var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Запускает какой бы то ни было сервер:)",
	Long: `Более длинное описание того, что делает команда.
		Несколько строк текста.
		Ага, третья строка`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("runserver called")

		flag, err := cmd.Flags().GetString("test")

		if flag != "" {
			fmt.Println("runserver called with flag: ", flag)
		}
		return err
	},
}

func init() {
	RootCmd.AddCommand(runserverCmd)
	runserverCmd.Flags().String("test", "", "just test flag")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
