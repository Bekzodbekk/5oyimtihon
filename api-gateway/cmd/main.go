package main

import (
	"api-gateway/internal/http"
	bookingservice "api-gateway/internal/pkg/booking-service"
	hotelservice "api-gateway/internal/pkg/hotel-service"
	"api-gateway/internal/pkg/load"
	notifservice "api-gateway/internal/pkg/notif-service"
	userservice "api-gateway/internal/pkg/user-service"
	"api-gateway/internal/service"
	"fmt"
	"log"
)

func main() {
	cfg, err := load.LOAD("./config/config.yaml")
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

	notifClient, err := notifservice.DialWithNotifService(*cfg)
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
		notifClient,
	)

	r := http.NewGin(*cli)

	target := fmt.Sprintf("%s:%d", cfg.ApiGatewayHost, cfg.ApiGatewayPort)
	if err := r.Run(target); err != nil {
		log.Fatal(err)
	}
}
