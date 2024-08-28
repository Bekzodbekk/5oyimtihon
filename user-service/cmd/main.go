package main

import (
	"log"
	mongosh "user-service/internal/pkg/Mongosh"
	redis "user-service/internal/pkg/Redis"
	"user-service/internal/pkg/load"
	userservice "user-service/internal/pkg/user-service"
	"user-service/internal/repository"
	"user-service/internal/service"
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

	rdb, err := redis.ConnectRedis(*conf)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepo(mongoConn.Collection, rdb, *conf)
	service := service.NewService(repo)
	run := userservice.NewRunSerivce(*service)
	
	log.Printf("User Service Running on :%d port", conf.UserServicePort)
	if err := run.RUN(*conf); err != nil{
		log.Fatal(err)
	}
}