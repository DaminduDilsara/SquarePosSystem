package main

import (
	"SquarePosSystem/configurations"
	"SquarePosSystem/internal/domain/clients/square_client"
	"SquarePosSystem/internal/domain/services/location_service"
	"SquarePosSystem/internal/transport/http"
	v1 "SquarePosSystem/internal/transport/http/controllers/v1"
	"log"
	"os"
)

func main() {

	sig := make(chan os.Signal, 0)

	conf := configurations.LoadConfigurations()

	squareClient := square_client.NewSquareClient(conf)

	locationService := location_service.NewLocationService(squareClient)
	controllerV1 := v1.NewControllerV1(locationService)

	http.InitServer(conf, controllerV1)

	select {
	case <-sig:
		log.Println("Application is shutting down..")

		os.Exit(0)
	}

}
