package main

import (
	"fmt"
	"log"
	"notif-service/internal/http"
	"notif-service/internal/pkg/load"
	"notif-service/internal/pkg/redis"
)

func main() {
	conf, err := load.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	rdb, err := redis.ConnectRedis(*conf)
	if err != nil {
		log.Fatal(err)
	}

	r := http.NewGin(rdb)

	target := fmt.Sprintf("%s:%d", conf.NotifHost, conf.NotifPort)
	if err := r.Run(target); err != nil {
		log.Fatal(err)
	}
}
