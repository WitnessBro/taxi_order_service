package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runserverCmd represents the runserver command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Запускает какой бы то ни было сервер:)",
	Long: `Более длинное описание того, что делает команда.
		Несколько строк текста.
		Ага, третья строка`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runserver called")

		flag, _ := cmd.Flags().GetString("test")

		if flag != "" {
			fmt.Println("runserver called with flag: ", flag)
		}
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)
	runserverCmd.Flags().String("test", "", "just test flag")
}
