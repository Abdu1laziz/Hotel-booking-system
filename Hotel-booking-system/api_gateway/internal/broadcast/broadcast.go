package broadcast1

import (
	"api-gateway/models"
	"api-gateway/pkg/kafka/producer"
	"api-gateway/pkg/protos/booking"
	"api-gateway/pkg/protos/hotel"
	"api-gateway/pkg/protos/user"
	redmet "api-gateway/pkg/redis/method"
	mail "api-gateway/utils/email"
	token "api-gateway/utils/jwt"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Adjust struct {
	U   user.UserClient
	R   *redmet.Redis
	B   booking.BookHotelClient
	H   hotel.HotelClient
	Ctx context.Context
}

func (a *Adjust) Register(req *models.RegisterUserRequest) error {
	if _, err := a.R.RegisterGet(req.Email); err != nil {
		if err := a.R.Register(req); err != nil {
			log.Println(err)
			return err
		}
		code := mail.Sent(req.Email)
		if err := a.R.VerifyCodeRequest(&models.VerifyRequest{Email: req.Email, Code: code}); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	return errors.New("this email already exists")
}

func (a *Adjust) Verify(req *models.VerifyRequest) error {
	res, err := a.R.VerifyCodeResponse(req)
	if err != nil {
		log.Println(err)
		return err
	}

	if res != req.Code {
		return errors.New("password or email doesn't match")
	}

	user, err := a.R.RegisterGet(req.Email)
	if err != nil {
		log.Println(err)
		return err
	}

	return a.CreateUser(user)
}

func (a *Adjust) CreateUser(req *models.RegisterUserRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	return producer.Producer("create", data)
}

func (a *Adjust) Login(req *models.LogInRequest) (map[string]string, error) {
	res, err := a.U.LogIn(a.Ctx, &user.LogInRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if res.Status {
		token, err := token.CreateToken(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		lastUserRes, err := a.U.LastInserted(a.Ctx, &user.LastInsertedUser{})
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return map[string]string{fmt.Sprintf("your account is created with this id %v", lastUserRes.Id): token}, nil
	}
	return nil, errors.New("password or email doesn't match or is missing")
}

func (a *Adjust) GetUser(req *models.GetUserRequest) (*models.GetUserResponse, error) {
	userData, err := a.R.GetUser(req)
	if err != nil {
		res, err := a.U.GetUser(a.Ctx, &user.GetUserRequest{Id: req.ID})
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if err := a.R.SetUser(&models.GetUserResponse{ID: res.Id, Username: res.Username, Age: res.Age, Email: res.Email, LogOut: res.Logout}); err != nil {
			log.Println(err)
		}
		return &models.GetUserResponse{ID: res.Id, Username: res.Username, Age: res.Age, Email: res.Email, LogOut: res.Logout}, nil
	}

	fmt.Printf("ID: %v, Username: %v, Age: %v, Email: %v, Logout: %v\n", userData.ID, userData.Username, userData.Age, userData.Email, userData.LogOut)
	return &models.GetUserResponse{ID: userData.ID, Username: userData.Username, Age: userData.Age, Email: userData.Email, LogOut: userData.LogOut}, nil
}

func (a *Adjust) UpdateUser(req *models.UpdateUserRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := producer.Producer("update", data); err != nil {
		log.Println(err)
		return err
	}

	res, err := a.U.GetUser(a.Ctx, &user.GetUserRequest{Id: req.ID})
	if err != nil {
		log.Println(err)
	}

	return a.R.SetUser(&models.GetUserResponse{ID: res.Id, Username: res.Username, Age: res.Age, Email: res.Email, LogOut: res.Logout})
}

func (a *Adjust) DeleteUser(req *models.GetUserRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	return producer.Producer("delete", data)
}

func (a *Adjust) Logout(req *models.GetUserRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	return producer.Producer("logout", data)
}

func (a *Adjust) CreateHotel(req *models.CreateHotelRequest) error {
	_, err := a.H.CreateHotel(a.Ctx, &hotel.CreateHotelRequest{Name: req.Name, Location: req.Location, Rating: req.Rating, Address: req.Address})
	return err
}

func (a *Adjust) GetHotel(req *models.GetHotelRequest) (*models.GetHotelResponse, error) {
	res, err := a.H.GetHotel(a.Ctx, &hotel.GetHotelRequest{Id: req.ID})
	if err != nil {
		return nil, err
	}

	var rooms []*models.UpdateRoomRequest
	for _, v := range res.Rooms {
		rooms = append(rooms, &models.UpdateRoomRequest{
			Available:     v.Available,
			RoomType:      v.RoomType,
			PricePerNight: v.PricePerNight,
			ID:            v.Id,
			HotelID:       v.HotelId,
		})
	}

	return &models.GetHotelResponse{ID: req.ID, Name: res.Name, Location: res.Location, Rating: res.Rating, Address: res.Address, Rooms: rooms}, nil
}

func (a *Adjust) GetHotels(req *models.GetsRequest) ([]*models.UpdateHotelRequest, error) {
	res, err := a.H.Gets(a.Ctx, &hotel.GetsRequest{})
	if err != nil {
		return nil, err
	}

	var hotels []*models.UpdateHotelRequest
	for _, v := range res.Hotels {
		hotels = append(hotels, &models.UpdateHotelRequest{
			ID:       v.Id,
			Name:     v.Name,
			Location: v.Location,
			Rating:   v.Rating,
			Address:  v.Address,
		})
	}
	return hotels, nil
}

func (a *Adjust) UpdateHotel(req *models.UpdateHotelRequest) error {
	_, err := a.H.Update(a.Ctx, &hotel.UpdateHotelRequest{Id: req.ID, Name: req.Name, Location: req.Location, Rating: req.Rating, Address: req.Address})
	return err
}

func (a *Adjust) DeleteHotel(req *models.GetHotelRequest) error {
	_, err := a.H.Delete(a.Ctx, &hotel.GetHotelRequest{Id: req.ID})
	return err
}

func (a *Adjust) CreateRoom(req *models.CreateRoomRequest) error {
	_, err := a.H.CreateRoom(a.Ctx, &hotel.CreateRoomRequest{HotelId: req.HotelID, RoomType: req.RoomType, PricePerNight: req.PricePerNight})
	return err
}

func (a *Adjust) GetRoom(req *models.GetRoomRequest) (*models.UpdateRoomRequest, error) {
	res, err := a.H.Get(a.Ctx, &hotel.GetroomRequest{HotelId: req.HotelID, Id: req.ID})
	if err != nil {
		return nil, err
	}
	return &models.UpdateRoomRequest{Available: res.Available, RoomType: res.RoomType, PricePerNight: res.PricePerNight, ID: res.Id, HotelID: res.HotelId}, nil
}

func (a *Adjust) GetRooms(req *models.GetRoomRequest) (*models.GetRoomResponse, error) {
	res, err := a.H.GetRooms(a.Ctx, &hotel.GetroomRequest{HotelId: req.HotelID, Id: req.ID})
	if err != nil {
		return nil, err
	}

	var rooms []*models.UpdateRoomRequest
	for _, v := range res.Rooms {
		rooms = append(rooms, &models.UpdateRoomRequest{
			ID:            v.Id,
			HotelID:       v.HotelId,
			Available:     v.Available,
			RoomType:      v.RoomType,
			PricePerNight: v.PricePerNight,
		})
	}
	return &models.GetRoomResponse{Rooms: rooms}, nil
}

func (a *Adjust) UpdateRoom(req *models.UpdateRoomRequest) error {
	_, err := a.H.UpdateRoom(a.Ctx, &hotel.UpdateRoomRequest{Available: req.Available, RoomType: req.RoomType, PricePerNight: req.PricePerNight, Id: req.ID, HotelId: req.HotelID})
	return err
}

func (a *Adjust) DeleteRoom(req *models.GetRoomRequest) error {
	_, err := a.H.DeleteRoom(a.Ctx, &hotel.GetroomRequest{HotelId: req.HotelID, Id: req.ID})
	return err
}

func (a *Adjust) CreateBooking(req *models.BookHotelRequest) (*models.GeneralResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := a.B.Create(a.Ctx, &booking.Bytes{All: data})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: res.Message}, nil
}

func (a *Adjust) GetBooking(req *models.GetUsersBookRequest) (*models.GetUsersBookResponse, error) {
	res, err := a.B.Get(a.Ctx, &booking.GetUsersBookRequest{Id: req.ID})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GetUsersBookResponse{
		ID:           res.Id,
		UserID:       res.UserID,
		HotelID:      res.HotelID,
		RoomID:       res.RoomId,
		RoomType:     res.RoomType,
		CheckInDate:  res.CheckInDate.AsTime(),
		CheckOutDate: res.CheckOutDate.AsTime(),
		TotalAmount:  res.TotalAmount,
		Status:       res.Status,
	}, nil
}

func (a *Adjust) UpdateBooking(req *models.BookHotelUpdateRequest) (*models.GeneralResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := a.B.Update(a.Ctx, &booking.Bytes{All: data})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: res.Message}, nil
}

func (a *Adjust) DeleteBooking(req *models.CancelRoomRequest) (*models.GeneralResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := a.B.Delete(a.Ctx, &booking.Bytes{All: data})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: res.Message}, nil
}

func (a *Adjust) CreateWaitinglist(req *models.CreateWaitingList) (*models.GeneralResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := a.B.CreateWaiting(a.Ctx, &booking.Bytes{All: data})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: res.Message}, nil
}

func (a *Adjust) GetWaiting(req *models.GetWaitinglistRequest) (*models.GetWaitinglistResponse, error) {
	res, err := a.B.GetWaitinglist(a.Ctx, &booking.GetWaitinglistRequest{Id: req.ID})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GetWaitinglistResponse{
		UserID:       res.UserId,
		UserEmail:    res.UserEmail,
		RoomType:     res.RoomType,
		HotelID:      res.HotelId,
		CheckInDate:  res.CheckInDate.AsTime(),
		CheckOutDate: res.CheckOutDate.AsTime(),
		Status:       res.Status,
		ID:           res.Id,
	}, nil
}

func (a *Adjust) UpdateWaiting(req *models.UpdateWaitingListRequest) (*models.GeneralResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := a.B.UpdateWaiting(a.Ctx, &booking.Bytes{All: data})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: res.Message}, nil
}

func (a *Adjust) DeleteWaiting(req *models.DeleteWaitingList) (*models.GeneralResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := a.B.DeleteWaiting(a.Ctx, &booking.Bytes{All: data})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: res.Message}, nil
}
