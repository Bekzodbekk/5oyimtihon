package bookingservice

import (
	"api-gateway/internal/pkg/load"
	"fmt"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithBookingService(cfg load.Config) (*booking.BookingServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.DialServiceConfig.BookingServiceHost, cfg.DialServiceConfig.BookingServicePort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := booking.NewBookingServiceClient(conn)
	return &client, nil
}
