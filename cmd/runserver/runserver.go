package runserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "runserver",
	Short: "Запускает какой бы то ни было сервер:)",
	Long: `Более длинное описание того, что делает команда.
		Несколько строк текста.
		Ага, третья строка`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if RunServer(); RunServer != nil {
			fmt.Println("Error on server starts")
		}
		return nil
	},
}

func helloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func RunServer() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
