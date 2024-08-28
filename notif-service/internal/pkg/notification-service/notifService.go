package notificationservice

import (
	"fmt"
	"net"
	"notif-service/internal/pkg/load"
	"notif-service/internal/service"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/notif"
	"google.golang.org/grpc"
)

type RunService struct {
	serv service.Service
}

func NewRunService(serv service.Service) *RunService {
	return &RunService{
		serv: serv,
	}
}

func (r *RunService) RUN(conf load.Config) error {
	target := fmt.Sprintf("%s:%d", conf.NotifHost, conf.NotifPort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	notif.RegisterNotifServiceServer(server, &r.serv)

	return server.Serve(listener)
}
