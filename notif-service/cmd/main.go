package main

import (
	"log"
	"notif-service/internal/pkg/load"
	notificationservice "notif-service/internal/pkg/notification-service"
	"notif-service/internal/repository"
	"notif-service/internal/service"
)

func main() {
	conf, err := load.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewNotifRepo()
	service := service.NewSerivce(&repo)
	run := notificationservice.NewRunService(*service)

	log.Printf("Notification service running on :%d port", conf.NotifPort)
	if err := run.RUN(*conf); err != nil {
		log.Fatal(err)
	}
}
