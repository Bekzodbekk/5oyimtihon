package repository

import (
	"context"
	"hotel-service/genproto/hotel"

	"go.mongodb.org/mongo-driver/mongo"
)

type HotelRepo struct {
	coll *mongo.Collection
}

func NewHotelRepo(coll *mongo.Collection) HotelRepository {
	return &HotelRepo{
		coll: coll,
	}
}

func (h *HotelRepo) CreateHotel(ctx context.Context, req *hotel.CreateHotelReq) (*hotel.CreateHotelResp, error) {
	return nil, nil
}
func (h *HotelRepo) UpdateHotel(ctx context.Context, req *hotel.UpdateHotelReq) (*hotel.UpdateHotelResp, error) {
	return nil, nil
}
func (h *HotelRepo) DeleteHotel(ctx context.Context, req *hotel.DeleteHotelReq) (*hotel.DeleteHotelResp, error) {
	return nil, nil
}
func (h *HotelRepo) GetHotels(ctx context.Context, req *hotel.Void) (*hotel.GetAllHotelsST, error) {
	return nil, nil
}
func (h *HotelRepo) GetHotelById(ctx context.Context, req *hotel.GetHotelByIdReq) (*hotel.GetHotelByIdResp, error) {
	return nil, nil
}
func (h *HotelRepo) GetHotelRoomsAvailability(ctx context.Context, req *hotel.HotelRoomsAvailabilityReq) (*hotel.HotelRoomsAvailabilityResp, error) {
	return nil, nil
}
