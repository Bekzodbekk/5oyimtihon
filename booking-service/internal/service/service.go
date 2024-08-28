package service

import (
	"booking-service/internal/repository"
	"context"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
)

type Service struct {
	*booking.UnimplementedBookingServiceServer
	bookingRepo repository.BookRepository
}

func NewService(book repository.BookRepository) *Service {
	return &Service{
		bookingRepo: book,
	}
}

func (s *Service) CreateBooking(ctx context.Context, req *booking.CreateBookingReq) (*booking.CreateBookingResp, error) {
	return s.bookingRepo.CreateBooking(ctx, req)
}
func (s *Service) UpdateBooking(ctx context.Context, req *booking.UpdateBookingReq) (*booking.UpdateBookingResp, error) {
	return s.bookingRepo.UpdateBooking(ctx, req)
}
func (s *Service) DeleteBooking(ctx context.Context, req *booking.DeleteBookingReq) (*booking.DeleteBookingResp, error) {
	return s.bookingRepo.DeleteBooking(ctx, req)
}
func (s *Service) GetBooking(ctx context.Context, req *booking.GetBookingReq) (*booking.GetBookingResp, error) {
	return s.bookingRepo.GetBooking(ctx, req)
}
func (s *Service) GetBookingByUserId(ctx context.Context, req *booking.GetBookingByUserIdReq) (*booking.GetBookingByUserIdResp, error) {
	return s.bookingRepo.GetBookingByUserId(ctx, req)
}
