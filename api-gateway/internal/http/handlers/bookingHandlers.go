package handlers

import (
	"net/http"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/booking"
	"github.com/gin-gonic/gin"
)

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
