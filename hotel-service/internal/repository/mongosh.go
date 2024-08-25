package repository

import (
	"context"
	pb "hotel-service/genproto/hotel"
	"time"

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
	Rooms      []Room             `bson:"rooms"`
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
		Rooms:      []Room{},
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

	now := time.Now()
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
				UpdatedAt: now.String(),
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

func (h *HotelRepo) GetHotels(ctx context.Context, req *pb.Void) (*pb.GetAllHotelsST, error) {
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
		hotels = append(hotels, convertHotelToPb(&hotel))
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
		Hotel:   convertHotelToPb(&hotel),
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
		if room.Availability {
			availableRooms = append(availableRooms, convertRoomToPb(&room))
		}
	}

	return &pb.HotelRoomsAvailabilityResp{
		Status:  true,
		Message: "Available rooms retrieved successfully",
		Rooms:   availableRooms,
	}, nil
}

func convertHotelToPb(hotel *Hotel) *pb.Hotel {
	return &pb.Hotel{
		HotelId:    hotel.HotelId.Hex(),
		Name:       hotel.Name,
		Location:   hotel.Location,
		Rating:     hotel.Rating,
		Address:    hotel.Address,
		RoomsCount: hotel.RoomsCount,
		Room:       convertRoomToPb(&hotel.Rooms[0]),
		Cud: &pb.CUD{
			CreatedAt: hotel.CreatedAt,
			UpdatedAt: hotel.UpdatedAt,
			DeletedAt: hotel.DeletedAt,
		},
	}
}

func convertRoomToPb(room *Room) *pb.Room {
	return &pb.Room{
		RoomId:        room.RoomId.Hex(),
		RoomType:      room.RoomType,
		PricePerNight: float32(room.PricePerNight),
		Availability:  room.Availability,
	}
}
