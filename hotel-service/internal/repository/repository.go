package repository

import (
	"context"
	"hotel-service/genproto/hotel"
)

type HotelRepository interface {
	CreateHotel(ctx context.Context, req *hotel.CreateHotelReq) (*hotel.CreateHotelResp, error)
	UpdateHotel(ctx context.Context, req *hotel.UpdateHotelReq) (*hotel.UpdateHotelResp, error)
	DeleteHotel(ctx context.Context, req *hotel.DeleteHotelReq) (*hotel.DeleteHotelResp, error)
	GetHotels(ctx context.Context, req *hotel.Void) (*hotel.GetAllHotelsST, error)
	GetHotelById(ctx context.Context, req *hotel.GetHotelByIdReq) (*hotel.GetHotelByIdResp, error)
	GetHotelRoomsAvailability(ctx context.Context, req *hotel.HotelRoomsAvailabilityReq) (*hotel.HotelRoomsAvailabilityResp, error)
}
