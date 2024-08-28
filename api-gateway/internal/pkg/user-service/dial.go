package userservice

import (
	"api-gateway/internal/pkg/load"
	"fmt"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithUserService(cfg load.Config) (*user.UserServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.DialServiceConfig.UserServiceHost, cfg.DialServiceConfig.UserServicePort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := user.NewUserServiceClient(conn)
	return &client, nil
}
