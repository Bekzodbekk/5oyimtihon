package handlers

import (
	"github.com/Bekzodbekk/5oyimtihonProto/genproto/hotel"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new hotel
// @Description Create a new hotel with the provided details
// @Tags hotels
// @Accept json
// @Produce json
// @Param hotel body CreateHotelReq true "Hotel information"
// @Success 200 {object} CreateHotelResp
// @Failure 400 {object} Error
// @Router /hotels [post]
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

// @Summary Update an existing hotel
// @Description Update an existing hotel's details
// @Tags hotels
// @Accept json
// @Produce json
// @Param hotel_id path string true "Hotel ID"
// @Param hotel body UpdateHotelReq true "Updated hotel information"
// @Success 200 {object} UpdateHotelResp
// @Failure 400 {object} Error
// @Router /hotels/{hotel_id} [put]
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

// @Summary Delete a hotel
// @Description Delete a hotel by its ID
// @Tags hotels
// @Produce json
// @Param hotel_id path string true "Hotel ID"
// @Success 200 {object} DeleteHotelResp
// @Failure 400 {object} Error
// @Router /hotels/{hotel_id} [delete]
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

// @Summary Get all hotels
// @Description Retrieve a list of all hotels
// @Tags hotels
// @Produce json
// @Success 200 {object} GetAllHotelsST
// @Failure 400 {object} Error
// @Router /hotels [get]
func (h *HandlerST) GetAllHotels(ctx *gin.Context) {
	resp, err := h.ClientRepository.GetAllHotels(ctx, &hotel.Void{})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Get a hotel by ID
// @Description Retrieve a specific hotel by its ID
// @Tags hotels
// @Produce json
// @Param hotel_id path string true "Hotel ID"
// @Success 200 {object} GetHotelByIdResp
// @Failure 400 {object} Error
// @Router /hotels/{hotel_id} [get]
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

// @Summary Get hotel rooms availability
// @Description Retrieve the availability of rooms for a specific hotel
// @Tags hotels
// @Produce json
// @Param hotel_id path string true "Hotel ID"
// @Success 200 {object} HotelRoomsAvailabilityResp
// @Failure 400 {object} Error
// @Router /hotels/{hotel_id}/rooms/availability [get]
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
