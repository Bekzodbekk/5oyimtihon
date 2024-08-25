package hotelservice

import (
	"fmt"
	"hotel-service/genproto/hotel"
	"hotel-service/internal/pkg/load"
	"hotel-service/internal/service"
	"net"

	"google.golang.org/grpc"
)

type Service struct {
	service service.Service
}

func RunService(service service.Service) *Service {
	return &Service{
		service: service,
	}
}

func (s *Service) RUN(cfg load.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.HotelServiceHost, cfg.HotelServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	hotel.RegisterHotelServiceServer(serv, &s.service)
	return serv.Serve(listener)
}
