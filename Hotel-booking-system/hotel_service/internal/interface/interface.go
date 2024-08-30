package interface17

import (
	"context"
	"hotel-service/models"
	"hotel-service/pkg/proto/hotel"
)

type Hotel interface{
	CreateHotel(ctx context.Context,req *models.CreateHotelRequest)(*models.GeneralResponse,error)
	GetHotel(ctx context.Context,req *models.GetHotelRequest)(*models.GetHotelResponse,error)
	Gets(ctx context.Context,req *models.GetsRequest)([]*models.UpdateHotelRequest,error)
	UpdateHotel(ctx context.Context,req *models.UpdateHotelRequest)(*models.GeneralResponse,error)
	DeleteHotel(ctx context.Context,req *models.GetHotelRequest)(*models.GeneralResponse,error)
	CreateRoom(ctx context.Context,req *models.CreateRoomRequest)(*models.GeneralResponse,error)
	GetRoom(ctx context.Context,req *models.GetRoomRequest)(*models.UpdateRoomRequest,error)
	GetRooms(ctx context.Context,req *models.GetRoomRequest)(*models.GetRoomResponse,error)
	UpdateRooms(ctx context.Context,req *models.UpdateRoomRequest)(*models.GeneralResponse,error)
	DeleteRoom(ctx context.Context,req *models.GetRoomRequest)(*models.GeneralResponse,error)
}

type Adjust interface{
	CreateHotel(ctx context.Context,req *hotel.CreateHotelRequest)(*hotel.GeneralResponse,error)
	GetHotel(ctx context.Context,req *hotel.GetHotelRequest)(*hotel.GetHotelResponse,error)
	Gets(ctx context.Context,req *hotel.GetsRequest)(*hotel.GetsResponse,error)
	UpdateHotel(ctx context.Context,req *hotel.UpdateHotelRequest)(*hotel.GeneralResponse,error)
	DeleteHotel(ctx context.Context,req *hotel.GetHotelRequest)(*hotel.GeneralResponse,error)
	CreateRoom(ctx context.Context,req *hotel.CreateRoomRequest)(*hotel.GeneralResponse,error)
	GetRoom(ctx context.Context,req *hotel.GetroomRequest)(*hotel.UpdateRoomRequest,error)
	GetRooms(ctx context.Context,req *hotel.GetroomRequest)(*hotel.GetroomResponse,error)
	UpdateRooms(ctx context.Context,req *hotel.UpdateRoomRequest)(*hotel.GeneralResponse,error)
	DeleteRoom(ctx context.Context,req *hotel.GetroomRequest)(*hotel.GeneralResponse,error)
}