package handlers

import (
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
	"github.com/gin-gonic/gin"
)

func (h *HandlerST) CreateHotel(ctx *gin.Context) {
	var hotel hotel.CreateHotelReq
	err := ctx.BindJSON(&hotel)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.ClientRepository.CreateHotel(ctx, &hotel)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerST) UpdateHotel(ctx *gin.Context) {
	hotelId := ctx.Param("hotel_id")
	var hotel hotel.UpdateHotelReq

	err := ctx.BindJSON(&hotel)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	hotel.HotelId = hotelId

	resp, err := h.ClientRepository.UpdateHotel(ctx, &hotel)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerST) DeleteHotel(ctx *gin.Context) {
	hotelId := ctx.Param("hotel_id")
	resp, err := h.ClientRepository.DeleteHotel(ctx, &hotel.DeleteHotelReq{
		Id: hotelId,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerST) GetAllHotels(ctx *gin.Context) {
	resp, err := h.ClientRepository.GetAllHotels(ctx, &hotel.Void{})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerST) GetHotelById(ctx *gin.Context) {
	hotelId := ctx.Param("hotel_id")

	resp, err := h.ClientRepository.GetHotelById(ctx, &hotel.GetHotelByIdReq{
		Id: hotelId,
	})

	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, resp)
}
func (h *HandlerST) GetHotelRoomsAvailability(ctx *gin.Context) {
	hotelId := ctx.Param("hotel_id")
	resp, err := h.ClientRepository.GetHotelRoomsAvailability(ctx, &hotel.HotelRoomsAvailabilityReq{
		Id: hotelId,
	})

	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
