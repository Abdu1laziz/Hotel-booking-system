package interfaceservices

import (
	interface17 "booking-service/internal/interface"
	"booking-service/models"
	"booking-service/pkg/protos/booking"
	"context"
)

type Database struct {
	D interface17.Booking
}

type AdjustDatabase struct {
	A interface17.BookingAdjust
}

func (u *Database) Create(ctx context.Context, req *models.BookHotelRequest, price float64) (*models.GeneralResponse, error) {
	return u.D.Create(ctx, req, price)
}
func (u *Database) Get(ctx context.Context, req *models.GetUsersBookRequest) (*models.GetUsersBookResponse, error) {
	return u.D.Get(ctx, req)
}
func (u *Database) Update(ctx context.Context, req *models.BookHotelUpdateRequest, price float64) (*models.GeneralResponse, error) {
	return u.D.Update(ctx, req, price)
}
func (u *Database) Cancel(ctx context.Context, req *models.CancelRoomRequest) (*models.GeneralResponse, error) {
	return u.D.Cancel(ctx, req)
}
func (u *Database) GetRoomInfo(ctx context.Context, req *models.GetRoomInfo) (*models.GetUsersBookResponse, error) {
	return u.D.GetRoomInfo(ctx, req)
}
func (u *Database) CreateW(ctx context.Context, req *models.CreateWaitingList) (*models.GeneralResponse, error) {
	return u.D.CreateW(ctx, req)
}
func (u *Database) GetW(ctx context.Context, req *models.GetWaitinglistRequest) (*models.GetWaitinglistResponse, error) {
	return u.D.GetW(ctx, req)
}
func (u *Database) UpdateW(ctx context.Context, req *models.UpdateWaitingListRequest) (*models.GeneralResponse, error) {
	return u.D.UpdateW(ctx, req)
}
func (u *Database) DeleteW(ctx context.Context, req *models.DeleteWaitingList) (*models.GeneralResponse, error) {
	return u.D.DeleteW(ctx, req)
}
func (u *AdjustDatabase) Create(ctx context.Context, req *booking.BookHotelRequest) (*booking.GeneralResponse, error) {
	return u.A.Create(ctx, req)
}
func (u *AdjustDatabase) Get(ctx context.Context, req *booking.GetUsersBookRequest) (*booking.GetUsersBookResponse, error) {
	return u.A.Get(ctx, req)
}
func (u *AdjustDatabase) Update(ctx context.Context, req *booking.BookHotelUpdateRequest) (*booking.GeneralResponse, error) {
	return u.A.Update(ctx, req)
}
func (u *AdjustDatabase) Cancel(ctx context.Context, req *booking.CancelROomRequest) (*booking.GeneralResponse, error) {
	return u.A.Cancel(ctx, req)
}
func (u *AdjustDatabase) CreateW(ctx context.Context, req *booking.CreateWaitingList) (*booking.GeneralResponse, error) {
	return u.A.CreateW(ctx, req)
}
func (u *AdjustDatabase) GetW(ctx context.Context, req *booking.GetWaitinglistRequest) (*booking.GetWaitinglistResponse, error) {
	return u.A.GetW(ctx, req)
}
func (u *AdjustDatabase) UpdateW(ctx context.Context, req *booking.UpdateWaitingListRequest) (*booking.GeneralResponse, error) {
	return u.A.UpdateW(ctx, req)
}
func (u *AdjustDatabase) DeleteW(ctx context.Context, req *booking.DeleteWaitingList) (*booking.GeneralResponse, error) {
	return u.A.DeleteW(ctx, req)
}
