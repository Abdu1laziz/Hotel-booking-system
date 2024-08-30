package services

import (
	"context"
	interface17 "hotel-service/internal/interface"
	"hotel-service/models"
	"hotel-service/pkg/proto/hotel"
)

type Database struct {
	S interface17.Hotel
}

type Adjust struct {
	A interface17.Adjust
}

func (u *Database) CreateHotel(ctx context.Context, req *models.CreateHotelRequest) (*models.GeneralResponse, error) {
	return u.S.CreateHotel(ctx, req)
}

func (u *Database) GetHotel(ctx context.Context, req *models.GetHotelRequest) (*models.GetHotelResponse, error) {
	return u.S.GetHotel(ctx, req)
}

func (u *Database) Gets(ctx context.Context, req *models.GetsRequest) ([]*models.UpdateHotelRequest, error) {
	return u.S.Gets(ctx, req)
}

func (u *Database) UpdateHotel(ctx context.Context, req *models.UpdateHotelRequest) (*models.GeneralResponse, error) {
	return u.S.UpdateHotel(ctx, req)
}

func (u *Database) DeleteHotel(ctx context.Context, req *models.GetHotelRequest) (*models.GeneralResponse, error) {
	return u.S.DeleteHotel(ctx, req)
}

func (u *Database) CreateRoom(ctx context.Context, req *models.CreateRoomRequest) (*models.GeneralResponse, error) {
	return u.S.CreateRoom(ctx, req)
}

func (u *Database) GetRoom(ctx context.Context, req *models.GetRoomRequest) (*models.UpdateRoomRequest, error) {
	return u.S.GetRoom(ctx, req)
}

func (u *Database) GetRooms(ctx context.Context, req *models.GetRoomRequest) (*models.GetRoomResponse, error) {
	return u.S.GetRooms(ctx, req)
}

func (u *Database) UpdateRooms(ctx context.Context, req *models.UpdateRoomRequest) (*models.GeneralResponse, error) {
	return u.S.UpdateRooms(ctx, req)
}

func (u *Database) DeleteRoom(ctx context.Context, req *models.GetRoomRequest) (*models.GeneralResponse, error) {
	return u.S.DeleteRoom(ctx, req)
}

func (a *Adjust) CreateHotel(ctx context.Context, req *hotel.CreateHotelRequest) (*hotel.GeneralResponse, error) {
	return a.A.CreateHotel(ctx, req)
}

func (a *Adjust) GetHotel(ctx context.Context, req *hotel.GetHotelRequest) (*hotel.GetHotelResponse, error) {
	return a.A.GetHotel(ctx, req)
}

func (a *Adjust) Gets(ctx context.Context, req *hotel.GetsRequest) (*hotel.GetsResponse, error) {
	return a.A.Gets(ctx, req)
}

func (a *Adjust) UpdateHotel(ctx context.Context, req *hotel.UpdateHotelRequest) (*hotel.GeneralResponse, error) {
	return a.A.UpdateHotel(ctx, req)
}

func (a *Adjust) DeleteHotel(ctx context.Context, req *hotel.GetHotelRequest) (*hotel.GeneralResponse, error) {
	return a.A.DeleteHotel(ctx, req)
}

func (a *Adjust) CreateRoom(ctx context.Context, req *hotel.CreateRoomRequest) (*hotel.GeneralResponse, error) {
	return a.A.CreateRoom(ctx, req)
}

func (a *Adjust) GetRoom(ctx context.Context, req *hotel.GetroomRequest) (*hotel.UpdateRoomRequest, error) {
	return a.A.GetRoom(ctx, req)
}

func (a *Adjust) GetRooms(ctx context.Context, req *hotel.GetroomRequest) (*hotel.GetroomResponse, error) {
	return a.A.GetRooms(ctx, req)
}

func (a *Adjust) UpdateRooms(ctx context.Context, req *hotel.UpdateRoomRequest) (*hotel.GeneralResponse, error) {
	return a.A.UpdateRooms(ctx, req)
}

func (a *Adjust) DeleteRoom(ctx context.Context, req *hotel.GetroomRequest) (*hotel.GeneralResponse, error) {
	return a.A.DeleteRoom(ctx, req)
}
