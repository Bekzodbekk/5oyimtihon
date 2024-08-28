package handlers

import (
	"api-gateway/internal/service"
)

type HandlerST struct {
	ClientRepository *service.ServiceRepository
}

func NewHandler(cli service.ServiceRepository) *HandlerST {
	return &HandlerST{
		ClientRepository: &cli,
	}
}
