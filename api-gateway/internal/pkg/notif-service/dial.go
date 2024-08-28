package notifservice

import (
	"api-gateway/internal/pkg/load"
	"fmt"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/notif"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithNotifService(cfg load.Config) (*notif.NotifServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.DialServiceConfig.NotifServiceHost, cfg.DialServiceConfig.NotifServicePort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := notif.NewNotifServiceClient(conn)
	return &client, nil
}
