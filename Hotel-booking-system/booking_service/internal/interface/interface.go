package interface17

import (
	"booking-service/models"
	"booking-service/pkg/protos/booking"
	"context"
)

type Booking interface {
	Create(ctx context.Context, req *models.BookHotelRequest, price float64) (*models.GeneralResponse, error)
	Get(ctx context.Context, req *models.GetUsersBookRequest) (*models.GetUsersBookResponse, error)
	GetRoomInfo(ctx context.Context, req *models.GetRoomInfo) (*models.GetUsersBookResponse, error)
	Update(ctx context.Context, req *models.BookHotelUpdateRequest, price float64) (*models.GeneralResponse, error)
	Cancel(ctx context.Context, req *models.CancelRoomRequest) (*models.GeneralResponse, error)
	CreateW(ctx context.Context, req *models.CreateWaitingList) (*models.GeneralResponse, error)
	GetW(ctx context.Context, req *models.GetWaitinglistRequest) (*models.GetWaitinglistResponse, error)
	UpdateW(ctx context.Context, req *models.UpdateWaitingListRequest) (*models.GeneralResponse, error)
	DeleteW(ctx context.Context, req *models.DeleteWaitingList) (*models.GeneralResponse, error)
}

type BookingAdjust interface {
	Create(ctx context.Context, req *booking.BookHotelRequest) (*booking.GeneralResponse, error)
	Get(ctx context.Context, req *booking.GetUsersBookRequest) (*booking.GetUsersBookResponse, error)
	Update(ctx context.Context, req *booking.BookHotelUpdateRequest) (*booking.GeneralResponse, error)
	Cancel(ctx context.Context, req *booking.CancelROomRequest) (*booking.GeneralResponse, error)
	CreateW(ctx context.Context, req *booking.CreateWaitingList) (*booking.GeneralResponse, error)
	GetW(ctx context.Context, req *booking.GetWaitinglistRequest) (*booking.GetWaitinglistResponse, error)
	UpdateW(ctx context.Context, req *booking.UpdateWaitingListRequest) (*booking.GeneralResponse, error)
	DeleteW(ctx context.Context, req *booking.DeleteWaitingList) (*booking.GeneralResponse, error)
}
