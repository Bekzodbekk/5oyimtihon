package repository

import (
	"context"
	"time"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepo struct {
	coll *mongo.Collection
}

func NewBookRepo(coll *mongo.Collection) BookRepository {
	return &BookRepo{
		coll: coll,
	}
}

func (b *BookRepo) CreateBooking(ctx context.Context, req *booking.CreateBookingReq) (*booking.CreateBookingResp, error) {
	bookingDoc := bson.M{
		"user_id":        req.UserId,
		"hotel_id":       req.HotelId,
		"room_type":      req.RoomType,
		"check_in_date":  req.CheckInDate,
		"check_out_date": req.CheckOutDate,
		"total_amount":   req.TotalAmount,
		"created_at":     time.Now(),
	}

	_, err := b.coll.InsertOne(ctx, bookingDoc)
	if err != nil {
		return nil, err
	}

	return &booking.CreateBookingResp{
		UserId:       req.UserId,
		HotelId:      req.HotelId,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
		Status:       true,
	}, nil
}

func (b *BookRepo) UpdateBooking(ctx context.Context, req *booking.UpdateBookingReq) (*booking.UpdateBookingResp, error) {
	objectID, err := primitive.ObjectIDFromHex(req.BookingId)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"check_in_date":  req.CheckInDate,
			"check_out_date": req.CheckOutDate,
			"total_amount":   req.TotalAmount,
			"updated_at":     time.Now(),
		},
	}

	_, err = b.coll.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	var updatedBooking bson.M
	err = b.coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedBooking)
	if err != nil {
		return nil, err
	}

	return &booking.UpdateBookingResp{
		BookingId:    objectID.Hex(),
		UserId:       updatedBooking["user_id"].(string),
		HotelId:      updatedBooking["hotel_id"].(string),
		RoomType:     updatedBooking["room_type"].(string),
		CheckInDate:  updatedBooking["check_in_date"].(string),
		CheckOutDate: updatedBooking["check_out_date"].(string),
		TotalAmount:  updatedBooking["total_amount"].(float32),
		Status:       true,
	}, nil
}

func (b *BookRepo) DeleteBooking(ctx context.Context, req *booking.DeleteBookingReq) (*booking.DeleteBookingResp, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	result, err := b.coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return &booking.DeleteBookingResp{
			Status:    false,
			Message:   "Booking not found",
			BookingId: req.Id,
		}, nil
	}

	return &booking.DeleteBookingResp{
		Status:    true,
		Message:   "Booking deleted successfully",
		BookingId: req.Id,
	}, nil
}

func (b *BookRepo) GetBooking(ctx context.Context, req *booking.GetBookingReq) (*booking.GetBookingResp, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	var bookingDoc bson.M
	err = b.coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&bookingDoc)
	if err != nil {
		return nil, err
	}

	return &booking.GetBookingResp{
		BookingId:    req.Id,
		UserId:       bookingDoc["user_id"].(string),
		HotelId:      bookingDoc["hotel_id"].(string),
		RoomType:     bookingDoc["room_type"].(string),
		CheckInDate:  bookingDoc["check_in_date"].(string),
		CheckOutDate: bookingDoc["check_out_date"].(string),
		TotalAmount:  bookingDoc["total_amount"].(float32),
		Status:       true,
	}, nil
}

func (b *BookRepo) GetBookingByUserId(ctx context.Context, req *booking.GetBookingByUserIdReq) (*booking.GetBookingByUserIdResp, error) {
	cursor, err := b.coll.Find(ctx, bson.M{"user_id": req.Id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var bookings []*booking.Booking
	for cursor.Next(ctx) {
		var bookingDoc bson.M
		if err := cursor.Decode(&bookingDoc); err != nil {
			return nil, err
		}

		booking := &booking.Booking{
			BookingId:    bookingDoc["_id"].(primitive.ObjectID).Hex(),
			UserId:       bookingDoc["user_id"].(string),
			HotelId:      bookingDoc["hotel_id"].(string),
			RoomType:     bookingDoc["room_type"].(string),
			CheckInDate:  bookingDoc["check_in_date"].(string),
			CheckOutDate: bookingDoc["check_out_date"].(string),
			TotalAmount:  bookingDoc["total_amount"].(float32),
		}
		bookings = append(bookings, booking)
	}

	return &booking.GetBookingByUserIdResp{
		Books: bookings,
	}, nil
}
