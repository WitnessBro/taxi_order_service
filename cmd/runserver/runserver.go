package runserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"taxi_order_service/config"
)

var Cmd = &cobra.Command{
	Use:   "runserver",
	Short: "Запускает какой бы то ни было сервер:)",
	Long: `Более длинное описание того, что делает команда.
		Несколько строк текста.
		Ага, третья строка`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := config.NewConfig()
		if err != nil {
			return fmt.Errorf("can’t read config: %w", err)
		}
		if err := RunServer(conf); err != nil {
			return fmt.Errorf("can’t run server: %w", err)
		}
		return nil
	},
}

func helloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func RunServer(config *config.Config) error {
	http.HandleFunc("/", helloWorldHandler)
	if err := http.ListenAndServe(config.Address, nil); err != nil {
		return fmt.Errorf("can't listen server on address %s: %w", config.Address, err)
	}
	return nil
}
