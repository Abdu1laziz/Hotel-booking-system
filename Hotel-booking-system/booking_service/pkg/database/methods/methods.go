package methods

import (
	"booking-service/models"
	sqlbuilder "booking-service/pkg/database/sql"
	"booking-service/pkg/protos/booking"
	"context"
	"database/sql"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Database struct {
	Db    *sql.DB
	Price float64
}

// Create(ctx context.Context, req *models.BookHotelRequest, price float64) (*models.GeneralResponse, error)
// Get(ctx context.Context, req *models.GetUsersBookRequest) (*models.GetUsersBookResponse, error)
// Update(ctx context.Context, req *models.BookHotelUpdateRequest) (*models.GeneralResponse, error)
// Cancel(ctx context.Context, req *models.CancelRoomRequest) (*models.GeneralResponse, error)

// type WaitingList interface{
// 	Create(ctx context.Context,req *models.CreateWaitingList)(*models.GeneralResponse,error)
// 	Get(ctx context.Context,req *models.GetWaitinglistRequest)(*models.GetWaitinglistResponse,error)
// 	UpdateW(ctx context.Context,req *models.UpdateWaitingListRequest)(*models.GeneralResponse,error)
// 	DeleteW(ctx context.Context,req *models.DeleteWaitingList)(*models.GeneralResponse,error)
// }

func (u *Database) Create(ctx context.Context, req *models.BookHotelRequest, price float64) (*models.GeneralResponse, error) {
	u.Price = price
	query, args, err := sqlbuilder.Create(req, price)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int

	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(id)
	return &models.GeneralResponse{Message: fmt.Sprintf("%v", id)}, nil
}

func (u *Database) Get(ctx context.Context, req *models.GetUsersBookRequest) (*models.GetUsersBookResponse, error) {
	query, args, err := sqlbuilder.Get(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var res models.GetUsersBookResponse

	err = u.Db.QueryRow(query, args...).Scan(
		&res.ID,
		&res.UserID,
		&res.HotelID,
		&res.RoomID,
		&res.RoomType,
		&res.CheckInDate,
		&res.CheckOutDate,
		&res.TotalAmount,
		&res.Status,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

func (u *Database) Update(ctx context.Context, req *models.BookHotelUpdateRequest, price float64) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.Update(req, price)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int

	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Booking is updating with this id %v", id)}, nil
}

func (u *Database) Cancel(ctx context.Context, req *models.CancelRoomRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.Cancel(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Booking is deleting with this id %v", req.ID)}, nil
}

func (u *Database) CreateW(ctx context.Context, req *models.CreateWaitingList) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.CreateWaitingList(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int

	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("User is adding to waiting list with this id %v", id)}, nil
}

func (u *Database) GetW(ctx context.Context, req *models.GetWaitinglistRequest) (*models.GetWaitinglistResponse, error) {
	query, args, err := sqlbuilder.GetWaitingList(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var res models.GetWaitinglistResponse

	if err := u.Db.QueryRow(query, args...).Scan(
		&res.ID,
		&res.UserID,
		&res.HotelID,
		&res.RoomType,
		&res.UserEmail,
		&res.CheckInDate,
		&res.CheckOutDate,
		&res.Status,
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

func (u *Database) UpdateW(ctx context.Context, req *models.UpdateWaitingListRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.UpdateWaitingList(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int

	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Waiting List  is updating with this id %v", id)}, nil
}

func (u *Database) DeleteW(ctx context.Context, req *models.DeleteWaitingList) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.DeleteWaitingList(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("waiting List  is deleting  this id %v", req.ID)}, nil
}

func (u *Database) GetRoomInfo(ctx context.Context, req *models.GetRoomInfo) (*models.GetUsersBookResponse, error) {
	query, args, err := sqlbuilder.GetRoomInfo(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var res models.GetUsersBookResponse

	err = u.Db.QueryRow(query, args...).Scan(
		&res.ID,
		&res.UserID,
		&res.HotelID,
		&res.RoomID,
		&res.RoomType,
		&res.CheckInDate,
		&res.CheckOutDate,
		&res.TotalAmount,
		&res.Status,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

func GetAllWAitingUSers() (*booking.Response, error) {
	db, err := sql.Open("postgres", "postgres://postgres:2005@localhost/booking?sslmode=disable")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	query, args, err := sqlbuilder.GetAllUsers()
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var Users []*booking.GetWaitinglistResponse

	for rows.Next() {
		var all models.GetWaitinglistResponse
		if err := rows.Scan(&all.ID, &all.UserID, &all.HotelID, &all.RoomType, &all.UserEmail, &all.CheckInDate, &all.CheckOutDate, &all.Status); err != nil {
			log.Println(err)
			return nil, err
		}
		var all1 = booking.GetWaitinglistResponse{
			UserId:       all.UserID,
			UserEmail:    all.UserEmail,
			RoomType:     all.RoomType,
			HotelId:      all.HotelID,
			CheckInDate:  timestamppb.New(all.CheckInDate),
			CheckOutDate: timestamppb.New(all.CheckOutDate),
			Status:       all.Status,
			Id:           all.ID}
		Users = append(Users, &all1)
	}
	return &booking.Response{Users: Users}, nil
}
