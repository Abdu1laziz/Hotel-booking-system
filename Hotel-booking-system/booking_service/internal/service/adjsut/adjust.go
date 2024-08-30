package adjsut

import (
	interfaceservices "booking-service/internal/interface/services"
	"booking-service/models"
	"booking-service/pkg/protos/booking"
	"booking-service/pkg/protos/hotel"
	notificationss "booking-service/pkg/protos/notification"
	"booking-service/pkg/protos/user"
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Adjust отвечает за обработку запросов на бронирование
type Adjust struct {
	User  user.UserClient
	Hotel hotel.HotelClient
	S     *interfaceservices.Database
	N     notificationss.NotificationClient
}

var userID int

// Create обрабатывает запрос на создание бронирования
func (u *Adjust) Create(ctx context.Context, req *booking.BookHotelRequest) (*booking.GeneralResponse, error) {
	email, err := u.CheckUser(ctx, req)
	if err != nil {
		return nil, err
	}

	_, err = u.CheckHotel(ctx, req)
	if err != nil {
		if errors.Is(err, models.ErrHotelNotFound) || errors.Is(err, models.ErrRoomNotFound) {
			_, notifyErr := u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: req.UserID, Message: err.Error()})
			if notifyErr != nil {
				log.Println(notifyErr)
			}
			return nil, err
		}

		if errors.Is(err, models.ErrRoomNotAvailable) {
			return u.handleWaitingList(ctx, req, email)
		}
	}

	return u.processBooking(ctx, req, email)
}

// handleWaitingList обрабатывает добавление в список ожидания
func (u *Adjust) handleWaitingList(ctx context.Context, req *booking.BookHotelRequest, email string) (*booking.GeneralResponse, error) {
	newReq := models.CreateWaitingList{
		UserID:       req.UserID,
		UserEmail:    email,
		RoomType:     req.RoomType,
		HotelID:      req.HotelID,
		CheckInDate:  req.CheckInDate.AsTime(),
		CheckOutDate: req.CheckOutDate.AsTime(),
	}
	res, err := u.S.CreateW(ctx, &newReq)
	if err != nil {
		log.Println("Waiting list error:", err)
		return nil, err
	}

	userID = int(newReq.UserID)
	_, notifyErr := u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: req.UserID, Message: res.Message})
	if notifyErr != nil {
		log.Println(notifyErr)
	}
	return &booking.GeneralResponse{Message: res.Message}, nil
}

// processBooking обрабатывает успешное бронирование
func (u *Adjust) processBooking(ctx context.Context, req *booking.BookHotelRequest, email string) (*booking.GeneralResponse, error) {
	res1, err := u.Hotel.Get(ctx, &hotel.GetroomRequest{HotelId: req.HotelID, Id: req.RoomId})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	newReq := models.BookHotelRequest{
		UserID:       req.UserID,
		RoomID:       req.RoomId,
		RoomType:     req.RoomType,
		HotelID:      req.HotelID,
		CheckInDate:  req.CheckInDate.AsTime(),
		CheckOutDate: req.CheckOutDate.AsTime(),
	}
	res, err := u.S.Create(ctx, &newReq, float64(res1.PricePerNight))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := u.updateRoomAvailability(ctx, req); err != nil {
		return nil, err
	}

	if err := u.sendNotifications(ctx, email, res.Message, req.UserID); err != nil {
		return nil, err
	}

	userID = int(newReq.UserID)
	return &booking.GeneralResponse{Message: res.Message}, nil
}

// updateRoomAvailability обновляет доступность номера
func (u *Adjust) updateRoomAvailability(ctx context.Context, req *booking.BookHotelRequest) error {
	_, err := u.Hotel.UpdateRoom(ctx, &hotel.UpdateRoomRequest{Available: false, HotelId: req.HotelID, Id: req.RoomId})
	return err
}

// sendNotifications отправляет уведомления пользователю
func (u *Adjust) sendNotifications(ctx context.Context, email, message string, userID int64) error {
	_, err := u.N.Email(ctx, &notificationss.EmailSend{Email: email, Message: fmt.Sprintf("Congratulations on successfully booking your room! Your booking ID is %v", message)})
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: userID, Message: fmt.Sprintf("Congratulations on successfully booking your room! Your booking ID is %v", message)})
	return err
}

// Get обрабатывает запрос на получение информации о бронировании
func (u *Adjust) Get(ctx context.Context, req *booking.GetUsersBookRequest) (*booking.GetUsersBookResponse, error) {
	res, err := u.S.Get(ctx, &models.GetUsersBookRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &booking.GetUsersBookResponse{
		Id:           res.ID,
		UserID:       res.UserID,
		HotelID:      res.HotelID,
		RoomId:       res.RoomID,
		RoomType:     res.RoomType,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		TotalAmount:  res.TotalAmount,
		Status:       res.Status,
	}, nil
}

// Update обрабатывает запрос на обновление бронирования
func (u *Adjust) Update(ctx context.Context, req *booking.BookHotelUpdateRequest) (*booking.GeneralResponse, error) {
	res1, err := u.Hotel.Get(ctx, &hotel.GetroomRequest{HotelId: int32(hotelid), Id: req.RoomId})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := u.S.Update(ctx, &models.BookHotelUpdateRequest{
		ID:           req.Id,
		RoomID:       req.RoomId,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
	}, float64(res1.PricePerNight))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: int32(userID), Message: "Your room info was successfully updated"})
	if err != nil {
		log.Println(err)
	}
	return &booking.GeneralResponse{Message: res.Message}, nil
}

// Cancel обрабатывает запрос на отмену бронирования
func (u *Adjust) Cancel(ctx context.Context, req *booking.CancelROomRequest) (*booking.GeneralResponse, error) {
	info, err := u.S.Get(ctx, &models.GetUsersBookRequest{ID: req.Id})
	if err != nil {
		log.Println("info", err)
		return nil, err
	}

	_, err = u.Hotel.UpdateRoom(ctx, &hotel.UpdateRoomRequest{Available: true, HotelId: info.HotelID, Id: info.RoomID})
	if err != nil {
		log.Println("Error updating room availability:", err)
		return nil, err
	}

	_, err = u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: int32(userID), Message: "You successfully cancelled the room"})
	if err != nil {
		log.Println("Notification error:", err)
	}

	res, err := u.S.Cancel(ctx, &models.CancelRoomRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &booking.GeneralResponse{Message: res.Message}, nil
}

// CreateW обрабатывает запрос на создание записи в ожидании
func (u *Adjust) CreateW(ctx context.Context, req *booking.CreateWaitingList) (*booking.GeneralResponse, error) {
	newReq := models.CreateWaitingList{
		UserID:       req.UserId,
		UserEmail:    req.UserEmail,
		RoomType:     req.RoomType,
		HotelID:      req.HotelId,
		CheckInDate:  req.CheckInDate.AsTime(),
		CheckOutDate: req.CheckOutDate.AsTime(),
	}
	res, err := u.S.CreateW(ctx, &newReq)
	if err != nil {
		log.Println("Waiting list error:", err)
		return nil, err
	}

	_, err = u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: int32(userID), Message: "You have been added to the waiting list"})
	if err != nil {
		log.Println(err)
	}
	return &booking.GeneralResponse{Message: res.Message}, nil
}

// GetW обрабатывает запрос на получение информации о записи в ожидании
func (u *Adjust) GetW(ctx context.Context, req *booking.GetWaitinglistRequest) (*booking.GetWaitinglistResponse, error) {
	res, err := u.S.GetW(ctx, &models.GetWaitinglistRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &booking.GetWaitinglistResponse{
		UserId:       res.UserID,
		UserEmail:    res.UserEmail,
		RoomType:     res.RoomType,
		HotelId:      res.HotelID,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		Status:       res.Status,
		Id:           res.ID,
	}, nil
}

// UpdateW обрабатывает запрос на обновление записи в ожидании
func (u *Adjust) UpdateW(ctx context.Context, req *booking.UpdateWaitingListRequest) (*booking.GeneralResponse, error) {
	res, err := u.S.UpdateW(ctx, &models.UpdateWaitingListRequest{
		UserID:       req.UserId,
		RoomType:     req.RoomType,
		HotelID:      req.HotelId,
		CheckInDate:  req.CheckInDate.AsTime(),
		CheckOutDate: req.CheckOutDate.AsTime(),
		ID:           req.Id,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &booking.GeneralResponse{Message: res.Message}, nil
}

// DeleteW обрабатывает запрос на удаление записи из ожидания
func (u *Adjust) DeleteW(ctx context.Context, req *booking.DeleteWaitingList) (*booking.GeneralResponse, error) {
	res, err := u.S.DeleteW(ctx, &models.DeleteWaitingList{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &booking.GeneralResponse{Message: res.Message}, nil
}

// CheckUser проверяет пользователя по ID
func (u *Adjust) CheckUser(ctx context.Context, req *booking.BookHotelRequest) (string, error) {
	res, err := u.User.GetUser(ctx, &user.GetUserRequest{Id: req.UserID})
	if err != nil {
		log.Println(err)
		_, notifyErr := u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: req.UserID, Message: "No such user with this ID"})
		if notifyErr != nil {
			log.Println(notifyErr)
		}
		return "", errors.New("no such user with this ID")
	}
	if res.Age < int32(18) {
		_, notifyErr := u.N.Notification(ctx, &notificationss.ProduceMessage{UserId: req.UserID, Message: "You must be older than 18 years to book a room"})
		if notifyErr != nil {
			log.Println(notifyErr)
		}
		return "", errors.New("you must be old enough to book a room")
	}
	return res.Email, nil
}

// CheckHotel проверяет наличие отеля и комнат
func (u *Adjust) CheckHotel(ctx context.Context, req *booking.BookHotelRequest) (float64, error) {
	res, err := u.Hotel.GetRooms(ctx, &hotel.GetroomRequest{HotelId: req.HotelID})
	if err != nil {
		log.Println(err)
		return 0, models.ErrHotelNotFound
	}

	checkInDate := req.CheckInDate.AsTime()
	var availableRooms []*hotel.UpdateRoomRequest
	for _, v := range res.Rooms {
		if v.RoomType == req.RoomType {
			roomInfo, err := u.S.GetRoomInfo(ctx, &models.GetRoomInfo{HotelID: req.HotelID})
			if err != nil {
				log.Println(err)
				return 0, err
			}
			if roomInfo.CheckOutDate.Before(checkInDate) || roomInfo.CheckOutDate.Equal(checkInDate) {
				availableRooms = append(availableRooms, v)
			}
		}
	}

	if len(availableRooms) > 0 {
		return float64(availableRooms[0].PricePerNight), nil
	}

	return 0, models.ErrRoomNotAvailable
}
