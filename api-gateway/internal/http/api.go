package http

import (
	handlers "api-gateway/internal/http/handlers"
	"api-gateway/internal/service"
	"api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewGin(cli service.ServiceRepository) *gin.Engine {
	r := gin.Default()

	hnd := handlers.NewHandler(cli)

	r.POST("/api/users/register", hnd.Register)
	r.POST("/api/users/verify", hnd.Verify)
	r.POST("/api/users/login", hnd.Login)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Hotel routes
		protected.POST("/hotels", hnd.CreateHotel)
		protected.PUT("/hotels/:hotel_id", hnd.UpdateHotel)
		protected.DELETE("/hotels/:hotel_id", hnd.DeleteHotel)
		protected.GET("/hotels", hnd.GetAllHotels)
		protected.GET("/hotels/:hotel_id", hnd.GetHotelById)
		protected.GET("/hotels/:hotel_id/rooms/availability", hnd.GetHotelRoomsAvailability)

		// Booking routes
		protected.POST("/bookings", hnd.CreateBooking)
		protected.PUT("/bookings/:booking_id", hnd.UpdateBooking)
		protected.DELETE("/bookings/:booking_id", hnd.DeleteBooking)
		protected.GET("/bookings/:booking_id", hnd.GetBooking)
		protected.GET("/users/:user_id/bookings", hnd.GetBookingByUserId)

		// User routes
		protected.PUT("/users/:user_id", hnd.UpdateUser)
		protected.DELETE("/users/:user_id", hnd.DeleteUser)
		protected.GET("/users/:user_id", hnd.GetUserById)
	}
	return r
}
