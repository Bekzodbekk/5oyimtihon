package service

import (
	"context"
	"hotel-service/genproto/hotel"
	"hotel-service/internal/repository"
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

func (h *Service) CreateHotel(ctx context.Context, req *hotel.CreateHotelReq) (*hotel.CreateHotelResp, error) {
	return h.HotelRepo.CreateHotel(ctx, req)
}
func (h *Service) UpdateHotel(ctx context.Context, req *hotel.UpdateHotelReq) (*hotel.UpdateHotelResp, error) {
	return h.HotelRepo.UpdateHotel(ctx, req)

}
func (h *Service) DeleteHotel(ctx context.Context, req *hotel.DeleteHotelReq) (*hotel.DeleteHotelResp, error) {
	return h.HotelRepo.DeleteHotel(ctx, req)

}
func (h *Service) GetHotels(ctx context.Context, req *hotel.Void) (*hotel.GetAllHotelsST, error) {
	return h.HotelRepo.GetHotels(ctx, req)

}
func (h *Service) GetHotelById(ctx context.Context, req *hotel.GetHotelByIdReq) (*hotel.GetHotelByIdResp, error) {
	return h.HotelRepo.GetHotelById(ctx, req)

}
func (h *Service) GetHotelRoomsAvailability(ctx context.Context, req *hotel.HotelRoomsAvailabilityReq) (*hotel.HotelRoomsAvailabilityResp, error) {
	return h.HotelRepo.GetHotelRoomsAvailability(ctx, req)

}
