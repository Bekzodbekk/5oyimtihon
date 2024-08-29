package http

import (
	handlers "api-gateway/internal/http/handlers"
	"api-gateway/internal/service"
	"api-gateway/middleware"
	"crypto/tls"
	"net/http"

	_ "api-gateway/docs"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:9000
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(cli service.ServiceRepository, rdb *redis.Client) *http.Server {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	hnd := handlers.NewHandler(cli, rdb)

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

		protected.POST("/rooms", hnd.CreateRoom)
		protected.PUT("/rooms/:hotel_id/:room_id", hnd.UpdateRoom)
		protected.DELETE("/rooms/:hotel_id/:room_id", hnd.DeleteRoom)

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

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      ":9000",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv

	// return r
}
