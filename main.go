package main

import (
	"SquarePosSystem/configurations"
	"SquarePosSystem/internal/domain/clients/square_client"
	"SquarePosSystem/internal/domain/services/location_service"
	"SquarePosSystem/internal/domain/services/order_service"
	"SquarePosSystem/internal/domain/services/payment_service"
	"SquarePosSystem/internal/transport/http"
	v1 "SquarePosSystem/internal/transport/http/controllers/v1"
	"log"
	"os"
)

func main() {

	sig := make(chan os.Signal, 0)

	// loading configurations
	conf := configurations.LoadConfigurations()

	// initializing square clients
	squareLocationClient := square_client.NewSquareLocationClient(conf)
	squareOrderClient := square_client.NewSquareOrderClient(conf)
	squarePaymentClient := square_client.NewSquarePaymentClient(conf)

	// initializing services
	locationService := location_service.NewLocationService(squareLocationClient)
	orderService := order_service.NewOrderService(squareOrderClient)
	paymentService := payment_service.NewPaymentService(squarePaymentClient)

	// initializing controller
	controllerV1 := v1.NewControllerV1(locationService, orderService, paymentService)

	// initializing web server
	http.InitServer(conf, controllerV1)

	select {
	case <-sig:
		log.Println("Application is shutting down..")
		os.Exit(0)
	}

}
