package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
