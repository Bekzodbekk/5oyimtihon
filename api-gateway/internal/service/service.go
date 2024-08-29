package service

import (
	"context"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/notif"
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
)

type ServiceRepository struct {
	BookingServiceClient      booking.BookingServiceClient
	HotelServiceClient        hotel.HotelServiceClient
	UserServiceClient         user.UserServiceClient
	NotificationServiceClient notif.NotifServiceClient
}

func NewServiceRepository(
	hotel *hotel.HotelServiceClient,
	booking *booking.BookingServiceClient,
	user *user.UserServiceClient,
) *ServiceRepository {
	return &ServiceRepository{
		HotelServiceClient:        *hotel,
		BookingServiceClient:      *booking,
		UserServiceClient:         *user,
	}
}

func (s *ServiceRepository) CreateHotel(ctx context.Context, req *hotel.CreateHotelReq) (*hotel.CreateHotelResp, error) {
	return s.HotelServiceClient.CreateHotel(ctx, req)
}
func (s *ServiceRepository) UpdateHotel(ctx context.Context, req *hotel.UpdateHotelReq) (*hotel.UpdateHotelResp, error) {
	return s.HotelServiceClient.UpdateHotel(ctx, req)
}
func (s *ServiceRepository) DeleteHotel(ctx context.Context, req *hotel.DeleteHotelReq) (*hotel.DeleteHotelResp, error) {
	return s.HotelServiceClient.DeleteHotel(ctx, req)
}
func (s *ServiceRepository) GetAllHotels(ctx context.Context, req *hotel.Void) (*hotel.GetAllHotelsST, error) {
	return s.HotelServiceClient.GetAllHotels(ctx, req)
}
func (s *ServiceRepository) GetHotelById(ctx context.Context, req *hotel.GetHotelByIdReq) (*hotel.GetHotelByIdResp, error) {
	return s.HotelServiceClient.GetHotelById(ctx, req)
}
func (s *ServiceRepository) GetHotelRoomsAvailability(ctx context.Context, req *hotel.HotelRoomsAvailabilityReq) (*hotel.HotelRoomsAvailabilityResp, error) {
	return s.HotelServiceClient.GetHotelRoomAvailability(ctx, req)
}
func (s *ServiceRepository) CreateRoom(ctx context.Context, req *hotel.CreateRoomReq) (*hotel.CreateRoomResp, error) {
	return s.HotelServiceClient.CreateRoom(ctx, req)
}
func (s *ServiceRepository) UpdateRoom(ctx context.Context, req *hotel.UpdateRoomReq) (*hotel.UpdateRoomResp, error) {
	return s.HotelServiceClient.UpdateRoom(ctx, req)
}
func (s *ServiceRepository) DeleteRoom(ctx context.Context, req *hotel.DeleteRoomReq) (*hotel.DeleteRoomResp, error) {
	return s.HotelServiceClient.DeleteRoom(ctx, req)
}



func (s *ServiceRepository) CreateBooking(ctx context.Context, req *booking.CreateBookingReq) (*booking.CreateBookingResp, error) {
	return s.BookingServiceClient.CreateBooking(ctx, req)
}
func (s *ServiceRepository) UpdateBooking(ctx context.Context, req *booking.UpdateBookingReq) (*booking.UpdateBookingResp, error) {
	return s.BookingServiceClient.UpdateBooking(ctx, req)
}
func (s *ServiceRepository) DeleteBooking(ctx context.Context, req *booking.DeleteBookingReq) (*booking.DeleteBookingResp, error) {
	return s.BookingServiceClient.DeleteBooking(ctx, req)
}
func (s *ServiceRepository) GetBooking(ctx context.Context, req *booking.GetBookingReq) (*booking.GetBookingResp, error) {
	return s.BookingServiceClient.GetBooking(ctx, req)
}
func (s *ServiceRepository) GetBookingByUserId(ctx context.Context, req *booking.GetBookingByUserIdReq) (*booking.GetBookingByUserIdResp, error) {
	return s.BookingServiceClient.GetBookingByUserId(ctx, req)
}

func (s *ServiceRepository) Register(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	return s.UserServiceClient.Register(ctx, req)
}
func (s *ServiceRepository) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	return s.UserServiceClient.Login(ctx, req)
}
func (s *ServiceRepository) Verify(ctx context.Context, req *user.VerifyReq) (*user.VerifyResp, error) {
	return s.UserServiceClient.Verify(ctx, req)
}
func (s *ServiceRepository) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	return s.UserServiceClient.UpdateUser(ctx, req)
}
func (s *ServiceRepository) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (*user.DeleteUserResp, error) {
	return s.UserServiceClient.DeleteUser(ctx, req)
}
func (s *ServiceRepository) GetUserById(ctx context.Context, req *user.GetUserByIdReq) (*user.GetUserByIdResp, error) {
	return s.UserServiceClient.GetUserById(ctx, req)
}

