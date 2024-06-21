package runserver

import (
	"fmt"
	"net/http"
)

func _helloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func RunServer() {
	http.HandleFunc("/", _helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
