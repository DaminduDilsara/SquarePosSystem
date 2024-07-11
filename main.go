package main

import (
	"SquarePosSystem/configurations"
	"SquarePosSystem/internal/transport/http"
	"log"
	"os"
)

func main() {

	sig := make(chan os.Signal, 0)

	conf := configurations.LoadConfigurations()
	
	http.InitServer(conf)

	select {
	case <-sig:
		log.Println("Application is shutting down..")

		os.Exit(0)
	}

}
