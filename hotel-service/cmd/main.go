package main

import (
	hotelservice "hotel-service/internal/pkg/hotel-service"
	"hotel-service/internal/pkg/load"
	"hotel-service/internal/pkg/mongosh"
	"hotel-service/internal/repository"
	"hotel-service/internal/service"
	"log"
)

func main() {
	conf, err := load.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	mongoshConnection, err := mongosh.InitDB(*conf)
	if err != nil {
		log.Fatal(err)
	}

	hotelRepo := repository.NewHotelRepo(&mongoshConnection.Collection)
	service := service.NewService(*hotelRepo)
	runServ := hotelservice.RunService(*service)

	log.Printf("Hotel Service running on :%d port", conf.HotelServicePort)
	if err := runServ.RUN(*conf); err != nil {
		log.Fatal(err)
	}
}
