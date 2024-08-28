package repository

import (
	"context"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
)

type HotelRepository interface {
	CreateHotel(ctx context.Context, req *hotel.CreateHotelReq) (*hotel.CreateHotelResp, error)
	UpdateHotel(ctx context.Context, req *hotel.UpdateHotelReq) (*hotel.UpdateHotelResp, error)
	DeleteHotel(ctx context.Context, req *hotel.DeleteHotelReq) (*hotel.DeleteHotelResp, error)
	GetHotels(ctx context.Context, req *hotel.Void) (*hotel.GetAllHotelsST, error)
	GetHotelById(ctx context.Context, req *hotel.GetHotelByIdReq) (*hotel.GetHotelByIdResp, error)
	GetHotelRoomsAvailability(ctx context.Context, req *hotel.HotelRoomsAvailabilityReq) (*hotel.HotelRoomsAvailabilityResp, error)

	CreateRoom(ctx context.Context, req *hotel.CreateRoomReq) (*hotel.CreateRoomResp, error)
	UpdateRoom(ctx context.Context, req *hotel.UpdateRoomReq) (*hotel.UpdateRoomResp, error)
	DeleteRoom(ctx context.Context, req *hotel.DeleteRoomReq) (*hotel.DeleteRoomResp, error)
}
