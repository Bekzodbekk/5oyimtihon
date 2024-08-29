package handlers

import (
	"api-gateway/internal/service"

	"github.com/redis/go-redis/v9"
)

type HandlerST struct {
	ClientRepository *service.ServiceRepository
	rdb              *redis.Client
}

func NewHandler(cli service.ServiceRepository, rdb *redis.Client) *HandlerST {
	return &HandlerST{
		ClientRepository: &cli,
		rdb:              rdb,
	}
}
