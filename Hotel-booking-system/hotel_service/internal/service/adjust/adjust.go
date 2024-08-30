package adjustservice

import (
	"context"
	"errors"
	"hotel-service/internal/interface/services"
	"hotel-service/models"
	"hotel-service/pkg/proto/hotel"
	"log"
)

type Adjust struct {
	S *services.Database
}

func (u *Adjust) CreateHotel(ctx context.Context, req *hotel.CreateHotelRequest) (*hotel.GeneralResponse, error) {
	if req.Address != "" && req.Location != "" && req.Name != "" && req.Rating != 0 {
		var newreq = models.CreateHotelRequest{
			Name:     req.Name,
			Location: req.Location,
			Rating:   req.Rating,
			Address:  req.Address,
		}
		res, err := u.S.CreateHotel(ctx, &newreq)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &hotel.GeneralResponse{Message: res.Message}, nil
	}
	return nil, errors.New("missing fields")

}
func (u *Adjust) GetHotel(ctx context.Context, req *hotel.GetHotelRequest) (*hotel.GetHotelResponse, error) {
	res, err := u.S.GetHotel(ctx, &models.GetHotelRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var rooms []*hotel.UpdateRoomRequest
	for _, v := range res.Rooms {
		var all = hotel.UpdateRoomRequest{
			Available:     v.Available,
			HotelId:       v.HotelID,
			RoomType:      v.RoomType,
			Id:            v.ID,
			PricePerNight: v.PricePerNight,
		}
		rooms = append(rooms, &all)
	}
	return &hotel.GetHotelResponse{Id: res.ID, Name: res.Name, Location: res.Location, Rating: res.Rating, Address: res.Address, Rooms: rooms}, nil
}

func (u *Adjust) Gets(ctx context.Context, req *hotel.GetsRequest) (*hotel.GetsResponse, error) {
	res, err := u.S.Gets(ctx, &models.GetsRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var hotels []*hotel.UpdateHotelRequest

	for _, v := range res {
		var all = hotel.UpdateHotelRequest{
			Id:       v.ID,
			Name:     v.Name,
			Location: v.Location,
			Rating:   v.Rating,
			Address:  v.Address,
		}
		hotels = append(hotels, &all)
	}
	return &hotel.GetsResponse{Hotels: hotels}, nil
}

func (u *Adjust) UpdateHotel(ctx context.Context, req *hotel.UpdateHotelRequest) (*hotel.GeneralResponse, error) {
		var newreq = models.UpdateHotelRequest{
			ID:       req.Id,
			Name:     req.Name,
			Location: req.Location,
			Rating:   req.Rating,
			Address:  req.Address,
		}
		res, err := u.S.UpdateHotel(ctx, &newreq)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &hotel.GeneralResponse{Message: res.Message}, nil
}

func (u *Adjust) DeleteHotel(ctx context.Context, req *hotel.GetHotelRequest) (*hotel.GeneralResponse, error) {
	res, err := u.S.DeleteHotel(ctx, &models.GetHotelRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.GeneralResponse{Message: res.Message}, nil
}

func (u *Adjust) CreateRoom(ctx context.Context, req *hotel.CreateRoomRequest) (*hotel.GeneralResponse, error) {
	if req.HotelId!=0 && req.RoomType!="" && req.PricePerNight!=0{
		var newrew = models.CreateRoomRequest{
			HotelID:       req.HotelId,
			RoomType:      req.RoomType,
			PricePerNight: req.PricePerNight,
		}
		res, err := u.S.CreateRoom(ctx, &newrew)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &hotel.GeneralResponse{Message: res.Message}, nil
	}
	return nil,errors.New("missing field")
}

func (u *Adjust) GetRoom(ctx context.Context, req *hotel.GetroomRequest) (*hotel.UpdateRoomRequest, error) {
	res, err := u.S.GetRoom(ctx, &models.GetRoomRequest{HotelID: req.HotelId, ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.UpdateRoomRequest{HotelId: res.HotelID, Id: res.ID, RoomType: res.RoomType, PricePerNight: res.PricePerNight, Available: res.Available}, nil
}

func (u *Adjust) GetRooms(ctx context.Context, req *hotel.GetroomRequest) (*hotel.GetroomResponse, error) {
	res, err := u.S.GetRooms(ctx, &models.GetRoomRequest{HotelID: req.HotelId})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var rooms []*hotel.UpdateRoomRequest

	for _, v := range res.Rooms {
		var all = hotel.UpdateRoomRequest{
			HotelId:       v.HotelID,
			Id:            v.ID,
			RoomType:      v.RoomType,
			PricePerNight: v.PricePerNight,
			Available:     v.Available,
		}
		rooms = append(rooms, &all)
	}
	return &hotel.GetroomResponse{Rooms: rooms}, nil
}
func (u *Adjust) UpdateRooms(ctx context.Context, req *hotel.UpdateRoomRequest) (*hotel.GeneralResponse, error) {
	var newreq = models.UpdateRoomRequest{
		HotelID:       req.HotelId,
		ID:            req.Id,
		RoomType:      req.RoomType,
		PricePerNight: req.PricePerNight,
		Available:     req.Available,
	}
	res, err := u.S.UpdateRooms(ctx, &newreq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.GeneralResponse{Message: res.Message}, nil
}

func (u *Adjust) DeleteRoom(ctx context.Context, req *hotel.GetroomRequest) (*hotel.GeneralResponse, error) {
	res, err := u.S.DeleteRoom(ctx, &models.GetRoomRequest{HotelID: req.HotelId, ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.GeneralResponse{Message: res.Message}, nil
}
