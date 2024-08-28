package repository

import (
	"context"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/notif"
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
)

type ServiceRepository interface {

	// Hotel Service
	CreateHotel(ctx context.Context, req *hotel.CreateHotelReq) (*hotel.CreateHotelResp, error)
	UpdateHotel(ctx context.Context, req *hotel.UpdateHotelReq) (*hotel.UpdateHotelResp, error)
	DeleteHotel(ctx context.Context, req *hotel.DeleteHotelReq) (*hotel.DeleteHotelResp, error)
	GetHotels(ctx context.Context, req *hotel.Void) (*hotel.GetAllHotelsST, error)
	GetHotelById(ctx context.Context, req *hotel.GetHotelByIdReq) (*hotel.GetHotelByIdResp, error)
	GetHotelRoomsAvailability(ctx context.Context, req *hotel.HotelRoomsAvailabilityReq) (*hotel.HotelRoomsAvailabilityResp, error)

	// Booking Service
	CreateBooking(ctx context.Context, req *booking.CreateBookingReq) (*booking.CreateBookingResp, error)
	UpdateBooking(ctx context.Context, req *booking.UpdateBookingReq) (*booking.UpdateBookingResp, error)
	DeleteBooking(ctx context.Context, req *booking.DeleteBookingReq) (*booking.DeleteBookingResp, error)
	GetBooking(ctx context.Context, req *booking.GetBookingReq) (*booking.GetBookingResp, error)
	GetBookingByUserId(ctx context.Context, req *booking.GetBookingByUserIdReq) (*booking.GetBookingByUserIdResp, error)

	// User Service
	Register(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error)
	Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error)
	Verify(ctx context.Context, req *user.VerifyReq) (*user.VerifyResp, error)
	UpdateUser(ctx context.Context, req *user.UpdateUserReq) (*user.UpdateUserResp, error)
	DeleteUser(ctx context.Context, req *user.DeleteUserReq) (*user.DeleteUserResp, error)
	GetUserById(ctx context.Context, req *user.GetUserByIdReq) (*user.GetUserByIdResp, error)

	// Notification Service
	Notification(ctx context.Context, req *notif.NotificationReq) (*notif.VoidNotif, error)
}
