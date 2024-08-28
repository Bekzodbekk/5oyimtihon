package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
	pb "github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Hotel struct {
	HotelId    primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Location   string             `bson:"location"`
	Rating     int64              `bson:"rating"`
	Address    string             `bson:"address"`
	RoomsCount int64              `bson:"rooms_count"`
	Rooms      bson.A             `bson:"rooms"`
	CreatedAt  string             `bson:"created_at"`
	UpdatedAt  string             `bson:"updated_at"`
	DeletedAt  int64              `bson:"deleted_at"`
}

type Room struct {
	RoomId        primitive.ObjectID `bson:"_id,omitempty"`
	RoomType      string             `bson:"room_type"`
	PricePerNight float64            `bson:"price_per_night"`
	Availability  bool               `bson:"availability"`
}

type HotelRepo struct {
	coll *mongo.Collection
}

func NewHotelRepo(coll *mongo.Collection) *HotelRepo {
	return &HotelRepo{
		coll: coll,
	}
}

func (h *HotelRepo) CreateHotel(ctx context.Context, req *pb.CreateHotelReq) (*pb.CreateHotelResp, error) {
	now := time.Now().String()
	hotel := Hotel{
		Name:       req.Name,
		Location:   req.Location,
		Rating:     req.Rating,
		Address:    req.Address,
		RoomsCount: 0,
		Rooms:      bson.A{},
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  0,
	}

	result, err := h.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)

	return &pb.CreateHotelResp{
		Status:  true,
		Message: "Hotel created successfully",
		Hotel: &pb.Hotel{
			HotelId:    insertedID.Hex(),
			Name:       hotel.Name,
			Location:   hotel.Location,
			Rating:     hotel.Rating,
			Address:    hotel.Address,
			RoomsCount: hotel.RoomsCount,
			Room:       &pb.Room{},
			Cud: &pb.CUD{
				CreatedAt: hotel.CreatedAt,
				UpdatedAt: hotel.UpdatedAt,
				DeletedAt: hotel.DeletedAt,
			},
		},
	}, nil

}

func (h *HotelRepo) UpdateHotel(ctx context.Context, req *pb.UpdateHotelReq) (*pb.UpdateHotelResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.HotelId)
	if err != nil {
		return nil, err
	}

	now := time.Now().String()
	update := bson.M{
		"$set": bson.M{
			"name":       req.Name,
			"location":   req.Location,
			"rating":     req.Rating,
			"address":    req.Address,
			"updated_at": now,
		},
	}
	filter := bson.M{"_id": hotelId, "deleted_at": 0}
	result, err := h.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return &pb.UpdateHotelResp{
			Status:  false,
			Message: "Hotel not found",
		}, nil
	}

	return &pb.UpdateHotelResp{
		Status:  true,
		Message: "Hotel updated successfully",
		Hotel: &pb.Hotel{
			HotelId:  req.HotelId,
			Name:     req.Name,
			Location: req.Location,
			Rating:   req.Rating,
			Address:  req.Address,
			Cud: &pb.CUD{
				UpdatedAt: now,
			},
		},
	}, nil
}

func (h *HotelRepo) DeleteHotel(ctx context.Context, req *pb.DeleteHotelReq) (*pb.DeleteHotelResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": hotelId, "deleted_at": 0}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now().String(),
			"deleted_at": time.Now().Unix(),
		},
	}
	result, err := h.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return &pb.DeleteHotelResp{
			Status:  false,
			Message: "Hotel not found",
		}, nil
	}

	return &pb.DeleteHotelResp{
		Status:  true,
		Message: "Hotel deleted successfully",
	}, nil
}

func (h *HotelRepo) GetAllHotels(ctx context.Context, req *pb.Void) (*pb.GetAllHotelsST, error) {
	filter := bson.M{"deleted_at": 0}
	cursor, err := h.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var hotels []*pb.Hotel
	for cursor.Next(ctx) {
		var hotel Hotel
		if err := cursor.Decode(&hotel); err != nil {
			return nil, err
		}
		pbHotel := &pb.Hotel{
			HotelId:    hotel.HotelId.Hex(),
			Name:       hotel.Name,
			Location:   hotel.Location,
			Rating:     hotel.Rating,
			Address:    hotel.Address,
			RoomsCount: hotel.RoomsCount,
			Room:       &pb.Room{},
			Cud: &pb.CUD{
				CreatedAt: hotel.CreatedAt,
				UpdatedAt: hotel.UpdatedAt,
				DeletedAt: hotel.DeletedAt,
			},
		}
		hotels = append(hotels, pbHotel)
	}

	return &pb.GetAllHotelsST{
		Status:      true,
		Message:     "Hotels retrieved successfully",
		HotelsCount: int64(len(hotels)),
		Hotels:      hotels,
	}, nil
}

func (h *HotelRepo) GetHotelById(ctx context.Context, req *pb.GetHotelByIdReq) (*pb.GetHotelByIdResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": hotelId, "deleted_at": 0}
	var hotel Hotel
	err = h.coll.FindOne(ctx, filter).Decode(&hotel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &pb.GetHotelByIdResp{
				Status:  false,
				Message: "Hotel not found",
			}, nil
		}
		return nil, err
	}

	return &pb.GetHotelByIdResp{
		Status:  true,
		Message: "Hotel retrieved successfully",
		Hotel: &pb.Hotel{
			HotelId:    req.Id,
			Name:       hotel.Name,
			Location:   hotel.Location,
			Rating:     hotel.Rating,
			Address:    hotel.Address,
			RoomsCount: hotel.RoomsCount,
			Room:       &pb.Room{},
			Cud: &pb.CUD{
				CreatedAt: hotel.CreatedAt,
				UpdatedAt: hotel.UpdatedAt,
				DeletedAt: hotel.DeletedAt,
			},
		},
	}, nil
}

func (h *HotelRepo) GetHotelRoomsAvailability(ctx context.Context, req *pb.HotelRoomsAvailabilityReq) (*pb.HotelRoomsAvailabilityResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": hotelId, "deleted_at": 0}
	var hotel Hotel
	err = h.coll.FindOne(ctx, filter).Decode(&hotel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &pb.HotelRoomsAvailabilityResp{
				Status:  false,
				Message: "Hotel not found",
			}, nil
		}
		return nil, err
	}

	var availableRooms []*pb.Room
	for _, room := range hotel.Rooms {
		roomMap, ok := room.(bson.M)
		if !ok {
			continue
		}

		availability, ok := roomMap["availability"].(bool)
		if !ok || !availability {
			continue
		}

		pbRoom := &pb.Room{
			RoomId:        roomMap["_id"].(primitive.ObjectID).Hex(),
			RoomType:      roomMap["room_type"].(string),
			PricePerNight: float32(roomMap["price_per_night"].(float64)),
			Availability:  availability,
		}
		availableRooms = append(availableRooms, pbRoom)
	}

	return &pb.HotelRoomsAvailabilityResp{
		Status:  true,
		Message: "Available rooms retrieved successfully",
		Rooms:   availableRooms,
	}, nil
}

func (h *HotelRepo) CreateRoom(ctx context.Context, req *hotel.CreateRoomReq) (*hotel.CreateRoomResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.HotelId)
	if err != nil {
		return nil, err
	}

	newRoom := Room{
		RoomId:        primitive.NewObjectID(),
		RoomType:      req.RoomType,
		PricePerNight: float64(req.PricePerNight),
		Availability:  true,
	}

	filter := bson.M{"_id": hotelId, "deleted_at": 0}
	update := bson.M{
		"$push": bson.M{"rooms": newRoom},
		"$inc":  bson.M{"rooms_count": 1},
		"$set":  bson.M{"updated_at": time.Now().String()},
	}

	result, err := h.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("hotel not found")
	}

	return &hotel.CreateRoomResp{
		Status:  true,
		Message: "Room created successfully",
		Room: &pb.Room{
			RoomId:        newRoom.RoomId.Hex(),
			RoomType:      newRoom.RoomType,
			PricePerNight: float32(newRoom.PricePerNight),
			Availability:  newRoom.Availability,
		},
	}, nil
}

func (h *HotelRepo) UpdateRoom(ctx context.Context, req *hotel.UpdateRoomReq) (*hotel.UpdateRoomResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.HotelId)
	if err != nil {
		return nil, err
	}

	roomId, err := primitive.ObjectIDFromHex(req.RoomId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":        hotelId,
		"deleted_at": 0,
		"rooms._id":  roomId,
	}

	update := bson.M{
		"$set": bson.M{
			"rooms.$.room_type":       req.RoomType,
			"rooms.$.price_per_night": req.PricePerNight,
			"updated_at":              time.Now().String(),
		},
	}

	result, err := h.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("hotel or room not found")
	}

	return &hotel.UpdateRoomResp{
		Status:  true,
		Message: "Room updated successfully",
	}, nil
}

func (h *HotelRepo) DeleteRoom(ctx context.Context, req *hotel.DeleteRoomReq) (*hotel.DeleteRoomResp, error) {
	hotelId, err := primitive.ObjectIDFromHex(req.HotelId)
	if err != nil {
		return nil, err
	}

	roomId, err := primitive.ObjectIDFromHex(req.RoomId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": hotelId, "deleted_at": 0}
	update := bson.M{
		"$pull": bson.M{"rooms": bson.M{"_id": roomId}},
		"$inc":  bson.M{"rooms_count": -1},
		"$set":  bson.M{"updated_at": time.Now().String()},
	}

	result, err := h.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("hotel not found")
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("room not found")
	}

	return &hotel.DeleteRoomResp{
		Status:  true,
		Message: "Room deleted successfully",
	}, nil
}
