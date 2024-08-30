package models

import (
	"errors"
	"time"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}

type VerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	ID int32 `json:"id"`
}

type LogInResponse struct {
	Status bool `json:"status"`
}

type LastInsertedUser struct{}

type GetUserResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	LogOut   bool   `json:"logout"`
}

type UpdateUserRequest struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BookHotelRequest struct {
	UserID       int32     `json:"userID"`
	HotelID      int32     `json:"hotelID"`
	RoomID       int32     `json:"room_id"`
	RoomType     string    `json:"roomType"`
	CheckInDate  time.Time `json:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate"`
}

type GetUsersBookRequest struct {
	ID int32 `json:"id"`
}

type GetRoomInfo struct {
	HotelID int32 `json:"hotelID"`
	RoomID  int32 `json:"room_id"`
}
type GetUsersBookResponse struct {
	ID           int32     `json:"id"`
	UserID       int32     `json:"userID"`
	HotelID      int32     `json:"hotelID"`
	RoomID       int32     `json:"room_id"`
	RoomType     string    `json:"roomType"`
	CheckInDate  time.Time `json:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate"`
	TotalAmount  float32   `json:"totalAmount"`
	Status       string    `json:"status"`
}

type BookHotelUpdateRequest struct {
	ID           int32     `json:"id"`
	RoomID       int32     `json:"room_id"`
	RoomType     string    `json:"roomType"`
	CheckInDate  time.Time `json:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate"`
}

type CancelRoomRequest struct {
	ID int32 `json:"id"`
}

type CreateWaitingList struct {
	UserID       int32     `json:"user_id"`
	UserEmail    string    `json:"user_email"`
	RoomType     string    `json:"room_type"`
	HotelID      int32     `json:"hotel_id"`
	CheckInDate  time.Time `json:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate"`
}

type GetWaitinglistRequest struct {
	ID int32 `json:"id"`
}

type GetWaitinglistResponse struct {
	UserID       int32     `json:"user_id"`
	UserEmail    string    `json:"user_email"`
	RoomType     string    `json:"room_type"`
	HotelID      int32     `json:"hotel_id"`
	CheckInDate  time.Time `json:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate"`
	Status       string    `json:"status"`
	ID           int32     `json:"id"`
}

type UpdateWaitingListRequest struct {
	UserID       int32     `json:"user_id"`
	RoomType     string    `json:"room_type"`
	HotelID      int32     `json:"hotel_id"`
	CheckInDate  time.Time `json:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate"`
	ID           int32     `json:"id"`
}

type DeleteWaitingList struct {
	ID int32 `json:"id"`
}

type CreateHotelRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Rating   int32  `json:"rating"`
	Address  string `json:"address"`
}
type GetsRequest struct {
}

type UpdateHotelRequest struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Rating   int32  `json:"rating"`
	Address  string `json:"address"`
}

type GetHotelRequest struct {
	ID int32 `json:"id"`
}

type GetHotelResponse struct {
	ID       int32                `json:"id"`
	Name     string               `json:"name"`
	Location string               `json:"location"`
	Rating   int32                `json:"rating"`
	Address  string               `json:"address"`
	Rooms    []*UpdateRoomRequest `json:"rooms"`
}

type CreateRoomRequest struct {
	HotelID       int32   `json:"hotel_id"`
	RoomType      string  `json:"room_type"`
	PricePerNight float32 `json:"price_per_night"`
}

type GetRoomRequest struct {
	HotelID int32 `json:"hotel_id"`
	ID      int32 `json:"id"`
}

type GetRoomsRequest struct {
	HotelID int32 `json:"hotel_id"`
}

type GetRoomResponse struct {
	Rooms []*UpdateRoomRequest `json:"rooms"`
}

type UpdateRoomRequest struct {
	Available     bool    `json:"available"`
	RoomType      string  `json:"room_type"`
	PricePerNight float32 `json:"price_per_night"`
	ID            int32   `json:"id"`
	HotelID       int32   `json:"hotel_id"`
}

var (
	ErrHotelNotFound    = errors.New("there is no such hotel with this id")
	ErrRoomNotFound     = errors.New("no room found matching the given criteria")
	ErrRoomNotAvailable = errors.New("room is not available for the requested dates")
)
