package service

import (
	"context"
	"hotel-service/internal/repository"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
)

type Service struct {
	hotel.UnimplementedHotelServiceServer
	HotelRepo repository.HotelRepo
}

func NewService(hotelRepo repository.HotelRepo) *Service {
	return &Service{
		HotelRepo: hotelRepo,
	}
}

func (s *Service) CreateHotel(ctx context.Context, req *hotel.CreateHotelReq) (*hotel.CreateHotelResp, error) {
	return s.HotelRepo.CreateHotel(ctx, req)
}
func (s *Service) UpdateHotel(ctx context.Context, req *hotel.UpdateHotelReq) (*hotel.UpdateHotelResp, error) {
	return s.HotelRepo.UpdateHotel(ctx, req)

}
func (s *Service) DeleteHotel(ctx context.Context, req *hotel.DeleteHotelReq) (*hotel.DeleteHotelResp, error) {
	return s.HotelRepo.DeleteHotel(ctx, req)

}
func (s *Service) GetAllHotels(ctx context.Context, req *hotel.Void) (*hotel.GetAllHotelsST, error) {
	return s.HotelRepo.GetAllHotels(ctx, req)

}
func (s *Service) GetHotelById(ctx context.Context, req *hotel.GetHotelByIdReq) (*hotel.GetHotelByIdResp, error) {
	return s.HotelRepo.GetHotelById(ctx, req)

}
func (s *Service) GetHotelRoomsAvailability(ctx context.Context, req *hotel.HotelRoomsAvailabilityReq) (*hotel.HotelRoomsAvailabilityResp, error) {
	return s.HotelRepo.GetHotelRoomsAvailability(ctx, req)
}

func (s *Service) CreateRoom(ctx context.Context, req *hotel.CreateRoomReq) (*hotel.CreateRoomResp, error) {
	return s.HotelRepo.CreateRoom(ctx, req)
}
func (s *Service) UpdateRoom(ctx context.Context, req *hotel.UpdateRoomReq) (*hotel.UpdateRoomResp, error) {
	return s.HotelRepo.UpdateRoom(ctx, req)
}
func (s *Service) DeleteRoom(ctx context.Context, req *hotel.DeleteRoomReq) (*hotel.DeleteRoomResp, error) {
	return s.HotelRepo.DeleteRoom(ctx, req)
}
