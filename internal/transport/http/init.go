package http

import (
	"SquarePosSystem/configurations"
	"SquarePosSystem/internal/transport/http/engines"
	"fmt"
	"log"
	"net/http"
	"time"
)

var engine http.Server

func InitServer(
	conf *configurations.Config,
) {
	engine = http.Server{
		Addr:         fmt.Sprintf(":%v", conf.AppConfig.AppPort),
		Handler:      engines.NewEngine().GetEngine(),
		WriteTimeout: time.Second * conf.AppConfig.WriteTimeout,
		ReadTimeout:  time.Second * conf.AppConfig.ReadTimeOut,
		IdleTimeout:  time.Second * conf.AppConfig.IdleTimeout,
	}

	go func() {
		if err := engine.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(fmt.Sprintf("Failed to start default web server : %v", err))
		}
	}()
	log.Println(fmt.Sprintf("Starting default web server under port : %v", conf.AppConfig.AppPort))
}
