package repository

import (
	"context"
	pb "hotel-service/genproto/hotel"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *HotelRepo) CreateHotel(ctx context.Context, req *pb.CreateHotelReq) (*pb.CreateHotelResp, error) {
	hotelReq := bson.M{
		"name":        req.Name,
		"location":    req.Location,
		"rating":      req.Rating,
		"address":     req.Address,
		"rooms_count": 0,
		"rooms":       bson.A{},
		"created_at":  time.Now().String(),
		"updated_at":  time.Now().String(),
		"deleted_at":  0,
	}

	result, err := h.coll.InsertOne(ctx, hotelReq)
	if err != nil {
		return nil, err
	}

	return &pb.CreateHotelResp{
		Status:  true,
		Message: "Hotel created successfully",
		Hotel: &pb.Hotel{
			HotelId:    result.InsertedID.(string),
			Name:       req.Name,
			Location:   req.Location,
			Rating:     req.Rating,
			Address:    req.Address,
			RoomsCount: 0,
			Room:       &pb.Room{},
			Cud: &pb.CUD{
				CreatedAt: time.Now().String(),
				UpdatedAt: time.Now().String(),
				DeletedAt: 0,
			},
		},
	}, nil
}
func (h *HotelRepo) UpdateHotel(ctx context.Context, req *pb.UpdateHotelReq) (*pb.UpdateHotelResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.)
	hotelData := bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"location":    req.Location,
			"rating":      req.Rating,
			"address":     req.Address,
			"rooms_count": 0,
			"rooms":       bson.A{},
			"updated_at":  time.Now().String(),
		},
	}
	
	return nil, nil
}
func (h *HotelRepo) DeleteHotel(ctx context.Context, req *pb.DeleteHotelReq) (*pb.DeleteHotelResp, error) {
	return nil, nil
}
func (h *HotelRepo) GetHotels(ctx context.Context, req *pb.Void) (*pb.GetAllHotelsST, error) {
	return nil, nil
}
func (h *HotelRepo) GetHotelById(ctx context.Context, req *pb.GetHotelByIdReq) (*pb.GetHotelByIdResp, error) {
	return nil, nil
}
func (h *HotelRepo) GetHotelRoomsAvailability(ctx context.Context, req *pb.HotelRoomsAvailabilityReq) (*pb.HotelRoomsAvailabilityResp, error) {
	return nil, nil
}
