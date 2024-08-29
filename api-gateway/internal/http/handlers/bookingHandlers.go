package handlers

import (
	"net/http"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new booking
// @Description Create a new booking for a user
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body CreateBookingReq true "Booking information"
// @Success 200 {object} CreateBookingResp
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /bookings [post]
func (h *HandlerST) CreateBooking(ctx *gin.Context) {
	var book booking.CreateBookingReq
	if err := ctx.BindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.ClientRepository.CreateBooking(ctx, &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update a booking
// @Description Update an existing booking
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking_id path string true "Booking ID"
// @Param booking body UpdateBookingReq true "Updated booking information"
// @Success 200 {object} UpdateBookingResp
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /bookings/{booking_id} [put]
func (h *HandlerST) UpdateBooking(ctx *gin.Context) {
	bookingID := ctx.Param("booking_id")
	var book booking.UpdateBookingReq
	if err := ctx.BindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.BookingId = bookingID

	resp, err := h.ClientRepository.UpdateBooking(ctx, &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete a booking
// @Description Delete an existing booking
// @Tags bookings
// @Produce json
// @Param booking_id path string true "Booking ID"
// @Success 200 {object} DeleteBookingResp
// @Failure 500 {object} Error
// @Router /bookings/{booking_id} [delete]
func (h *HandlerST) DeleteBooking(ctx *gin.Context) {
	id := ctx.Param("booking_id")
	resp, err := h.ClientRepository.DeleteBooking(ctx, &booking.DeleteBookingReq{
		Id: id,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get bookings by user ID
// @Description Retrieve all bookings for a specific user
// @Tags bookings
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} GetBookingByUserIdResp
// @Failure 500 {object} Error
// @Router /bookings/user/{user_id} [get]
func (h *HandlerST) GetBookingByUserId(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	resp, err := h.ClientRepository.GetBookingByUserId(ctx, &booking.GetBookingByUserIdReq{
		Id: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get a booking by ID
// @Description Retrieve a specific booking by its ID
// @Tags bookings
// @Produce json
// @Param booking_id path string true "Booking ID"
// @Success 200 {object} GetBookingResp
// @Failure 500 {object} Error
// @Router /bookings/{booking_id} [get]
func (h *HandlerST) GetBooking(ctx *gin.Context) {
	bookingID := ctx.Param("booking_id")
	resp, err := h.ClientRepository.GetBooking(ctx, &booking.GetBookingReq{
		Id: bookingID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
