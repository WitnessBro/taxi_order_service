package runserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"taxi_order_service/config"
	"taxi_order_service/server"
	"taxi_order_service/server/locations"
	"taxi_order_service/services"
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
		locationService := services.NewLocationService([]string{"192.168.1.8:9092"}, "points")
		if err := RunServer(conf, locationService); err != nil {
			return fmt.Errorf("can’t run server: %w", err)
		}
		return nil
	},
}

func RunServer(config *config.Config, locationService *services.LocationService) error {
	locationHandler := locations.Handler{
		LocationService: locationService,
	}
	routerBuilder := server.RouterBuilder{
		LocationHandler: &locationHandler,
	}
	handler := routerBuilder.Build()
	if err := http.ListenAndServe(config.Address, handler); err != nil {
		return fmt.Errorf("can't listen server on address %s: %w", config.Address, err)
	}
	return nil
}
