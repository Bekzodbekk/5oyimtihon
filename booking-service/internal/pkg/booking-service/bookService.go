package userservice

import (
	"booking-service/internal/pkg/load"
	"booking-service/internal/service"
	"fmt"
	"net"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
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
	target := fmt.Sprintf("%s:%d", conf.BookingServiceHost, conf.BookingServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	booking.RegisterBookingServiceServer(server, &r.service)

	return server.Serve(listener)
}
