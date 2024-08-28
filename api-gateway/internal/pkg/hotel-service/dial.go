package hotelservice

import (
	"api-gateway/internal/pkg/load"
	"fmt"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithHotelService(cfg load.Config) (*hotel.HotelServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.DialServiceConfig.HotelServiceHost, cfg.DialServiceConfig.HotelServicePort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := hotel.NewHotelServiceClient(conn)
	return &client, nil
}
