package models

import (
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

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
	ID           int32                  `json:"id"`
	RoomID       int32                  `json:"room_id"`
	RoomType     string                 `json:"roomType"`
	CheckInDate  *timestamppb.Timestamp `json:"checkInDate"`
	CheckOutDate *timestamppb.Timestamp `json:"checkOutDate"`
}

type GeneralResponse struct {
	Message string `json:"message"`
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

var (
	ErrHotelNotFound    = errors.New("there is no such hotel with this id")
	ErrRoomNotFound     = errors.New("no room found matching the given criteria")
	ErrRoomNotAvailable = errors.New("room is not available for the requested dates")
)
