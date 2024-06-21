package main

import (
	"taxi_order_service/cmd"
	"taxi_order_service/cmd/runserver"
)

func main() {
	cmd.Execute()
	runserver.RunServer()
}
