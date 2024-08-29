package handlers

type User struct {
	UserID   string  `json:"user_id"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	CUD      CUDUser `json:"cud"`
}

type CUDUser struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

// ===============================================
type CreateUserReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// =================================================
type UpdateUserReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

// ===================================================
type DeleteUserReq struct {
	ID string `json:"id"`
}

type DeleteUserResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ================================================
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	Token     string `json:"token"`
	ExpiresIn string `json:"expires_in"`
}

// ====================================================
type GetUserByIdReq struct {
	ID string `json:"id"`
}

type GetUserByIdResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

// ==================================================
type VerifyReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Hotel struct {
	HotelID    string `json:"hotel_id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Rating     int64  `json:"rating"`
	Address    string `json:"address"`
	RoomsCount int64  `json:"rooms_count"`
	Room       Room   `json:"room"`
	CUD        CUD    `json:"cud"`
}

type Room struct {
	HotelID       string  `json:"hotel_id"`
	RoomID        string  `json:"room_id"`
	RoomType      string  `json:"room_type"`
	PricePerNight float32 `json:"price_per_night"`
	Availability  bool    `json:"availability"`
}

type CUD struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

// ========================================================================

type CreateRoomReq struct {
	HotelID       string  `json:"hotel_id"`
	RoomType      string  `json:"room_type"`
	PricePerNight float32 `json:"price_per_night"`
}

type CreateRoomResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Room    Room   `json:"room"`
}

type UpdateRoomReq struct {
	HotelID       string  `json:"hotel_id"`
	RoomID        string  `json:"room_id"`
	RoomType      string  `json:"room_type"`
	PricePerNight float32 `json:"price_per_night"`
}

type UpdateRoomResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type DeleteRoomReq struct {
	HotelID string `json:"hotel_id"`
	RoomID  string `json:"room_id"`
}

type DeleteRoomResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ========================================================================

type CreateHotelReq struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Rating   int64  `json:"rating"`
	Address  string `json:"address"`
}

type CreateHotelResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Hotel   Hotel  `json:"hotel"`
}

// ========================================================================

type UpdateHotelReq struct {
	HotelID  string `json:"hotel_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Rating   int64  `json:"rating"`
	Address  string `json:"address"`
}

type UpdateHotelResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Hotel   Hotel  `json:"hotel"`
}

// ============================================================================

type DeleteHotelReq struct {
	ID string `json:"id"`
}

type DeleteHotelResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==========================================================================

type Void struct{}

type GetAllHotelsST struct {
	Status      bool    `json:"status"`
	Message     string  `json:"message"`
	HotelsCount int64   `json:"hotels_count"`
	Hotels      []Hotel `json:"hotels"`
}

// =========================================================================

type GetHotelByIdReq struct {
	ID string `json:"id"`
}

type GetHotelByIdResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Hotel   Hotel  `json:"hotel"`
}

// ==========================================================================

type HotelRoomsAvailabilityReq struct {
	ID string `json:"id"`
}

type HotelRoomsAvailabilityResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Rooms   []Room `json:"rooms"`
}

type Booking struct {
	BookingID    string  `json:"booking_id"`
	UserID       string  `json:"user_id"`
	HotelID      string  `json:"hotel_id"`
	RoomType     string  `json:"room_type"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	TotalAmount  float32 `json:"total_amount"`
}

//  =========================================

type CreateBookingReq struct {
	UserID       string  `json:"user_id"`
	HotelID      string  `json:"hotel_id"`
	RoomType     string  `json:"room_type"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	TotalAmount  float32 `json:"total_amount"`
}

type CreateBookingResp struct {
	UserID       string  `json:"user_id"`
	HotelID      string  `json:"hotel_id"`
	RoomType     string  `json:"room_type"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	TotalAmount  float32 `json:"total_amount"`
	Status       bool    `json:"status"`
}

// =================================================

type GetBookingReq struct {
	ID string `json:"id"`
}

type GetBookingResp struct {
	BookingID    string  `json:"booking_id"`
	UserID       string  `json:"user_id"`
	HotelID      string  `json:"hotel_id"`
	RoomType     string  `json:"room_type"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	TotalAmount  float32 `json:"total_amount"`
	Status       bool    `json:"status"`
}

// ===================================================

type UpdateBookingReq struct {
	BookingID    string  `json:"booking_id"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	TotalAmount  float32 `json:"total_amount"`
}

type UpdateBookingResp struct {
	BookingID    string  `json:"booking_id"`
	UserID       string  `json:"user_id"`
	HotelID      string  `json:"hotel_id"`
	RoomType     string  `json:"room_type"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	TotalAmount  float32 `json:"total_amount"`
	Status       bool    `json:"status"`
}

// ==================================================

type DeleteBookingReq struct {
	ID string `json:"id"`
}

type DeleteBookingResp struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	BookingID string `json:"booking_id"`
}

// =================================================

type GetBookingByUserIdReq struct {
	ID string `json:"id"`
}

type GetBookingByUserIdResp struct {
	Books []Booking `json:"books"`
}
