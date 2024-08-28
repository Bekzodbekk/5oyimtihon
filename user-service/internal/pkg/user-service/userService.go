package userservice

import (
	"fmt"
	"net"
	"user-service/internal/pkg/load"
	"user-service/internal/service"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
	"google.golang.org/grpc"
)

type RunService struct {
	service service.Service
}

func NewRunSerivce(service service.Service) *RunService {
	return &RunService{
		service: service,
	}
}

func (r RunService) RUN(conf load.Config) error {
	target := fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &r.service)

	return server.Serve(listener)
}
