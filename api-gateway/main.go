package main

import (
	"api-gateway/internal/http"
	bookingservice "api-gateway/internal/pkg/booking-service"
	hotelservice "api-gateway/internal/pkg/hotel-service"
	"api-gateway/internal/pkg/load"
	"api-gateway/internal/pkg/redis"
	userservice "api-gateway/internal/pkg/user-service"
	"api-gateway/internal/service"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := load.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	rdb, err := redis.ConnectRedis(*cfg)
	if err != nil {
		log.Fatal(err)
	}


	hotelClient, err := hotelservice.DialWithHotelService(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	bookingClient, err := bookingservice.DialWithBookingService(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	userClient, err := userservice.DialWithUserService(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	cli := service.NewServiceRepository(
		hotelClient,
		bookingClient,
		userClient,
	)

	r := http.NewGin(*cli, rdb)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	addr := fmt.Sprintf(":%d", cfg.ApiGatewayPort)
	go func() {
		log.Printf("Starting API Gateway on: %s", addr)
		if err := r.ListenAndServeTLS(cfg.Cert, cfg.Key); err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Starting API Gateway on address: %s", addr)

	signalReceived := <-sigChan
	log.Printf("Received signal: %s", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := r.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}
	log.Printf("Graceful shutdown complete.")

}
