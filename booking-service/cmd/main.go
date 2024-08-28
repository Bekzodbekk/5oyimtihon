package main

import (
	mongosh "booking-service/internal/pkg/Mongosh"
	userservice "booking-service/internal/pkg/booking-service"
	"booking-service/internal/pkg/load"
	"booking-service/internal/repository"
	"booking-service/internal/service"
	"log"
)

func main() {
	conf, err := load.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	mongoConn, err := mongosh.InitDB(*conf)
	if err != nil {
		log.Fatal(err)
	}


	repo := repository.NewBookRepo(mongoConn.Collection)
	service := service.NewService(repo)
	run := userservice.NewRunSerivce(*service)

	log.Printf("Booking Service Running on :%d port", conf.BookingServicePort)
	if err := run.RUN(*conf); err != nil {
		log.Fatal(err)
	}
}
