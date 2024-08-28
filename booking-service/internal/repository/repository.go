package repository

import (
	"context"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
)

type BookRepository interface {
	CreateBooking(ctx context.Context, req *booking.CreateBookingReq) (*booking.CreateBookingResp, error)
	UpdateBooking(ctx context.Context, req *booking.UpdateBookingReq) (*booking.UpdateBookingResp, error)
	DeleteBooking(ctx context.Context, req *booking.DeleteBookingReq) (*booking.DeleteBookingResp, error)
	GetBooking(ctx context.Context, req *booking.GetBookingReq) (*booking.GetBookingResp, error)
	GetBookingByUserId(ctx context.Context, req *booking.GetBookingByUserIdReq) (*booking.GetBookingByUserIdResp, error)
}
