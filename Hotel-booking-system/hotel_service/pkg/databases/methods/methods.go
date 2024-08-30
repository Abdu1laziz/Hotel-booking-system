package methods

import (
	"context"
	"database/sql"
	"fmt"
	sqlbuilder "hotel-service/pkg/databases/sql"
	"hotel-service/models"
	"log"
)

type Database struct {
	Db *sql.DB
}

func (u *Database) CreateHotel(ctx context.Context, req *models.CreateHotelRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.CreateHotel(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int

	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Hotel has been created with this id %v", id)}, nil
}

func (u *Database) GetHotel(ctx context.Context, req *models.GetHotelRequest) (*models.GetHotelResponse, error) {
	query, args, err := sqlbuilder.GetHotel(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	q, a, err := sqlbuilder.GetRoomForHotel(int(req.ID))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := u.Db.Query(q, a...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var rooms []*models.UpdateRoomRequest

	for rows.Next() {
		var all models.UpdateRoomRequest
		if err := rows.Scan(&all.HotelID, &all.ID, &all.RoomType, &all.PricePerNight, &all.Available); err != nil {
			log.Println(err)
			return nil, err
		}
		rooms = append(rooms, &all)
	}

	var res models.GetHotelResponse

	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Name, &res.Location, &res.Rating, &res.Address); err != nil {
		log.Println(err)
		return nil, err
	}
	res.Rooms = rooms
	return &res, nil
}

func (u *Database) Gets(ctx context.Context, req *models.GetsRequest) ([]*models.UpdateHotelRequest, error) {
	query, args, err := sqlbuilder.GetsHotel(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res []*models.UpdateHotelRequest

	rows, err := u.Db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var all models.UpdateHotelRequest
		if err := rows.Scan(&all.ID, &all.Name, &all.Location, &all.Rating,&all.Address); err != nil {
			log.Println(err)
			return nil, err
		}
		res = append(res, &all)
	}
	return res, nil
}

func (u *Database) UpdateHotel(ctx context.Context, req *models.UpdateHotelRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.UpdateHotel(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(req)
	var id int

	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println("here is the error",err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Hotel has been updated with this id %v", id)}, nil
}

func (u *Database) DeleteHotel(ctx context.Context, req *models.GetHotelRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.DeleteHotel(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Hotel has been deleted with this id %v", req.ID)}, nil
}

func (u *Database) CreateRoom(ctx context.Context, req *models.CreateRoomRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.CreateRoom(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int
	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Room has been added with this id %v", id)}, nil
}

func (u *Database) GetRoom(ctx context.Context, req *models.GetRoomRequest) (*models.UpdateRoomRequest, error) {
	query, args, err := sqlbuilder.GetRoom(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res models.UpdateRoomRequest

	if err := u.Db.QueryRow(query, args...).Scan(&res.HotelID, &res.ID, &res.RoomType, &res.PricePerNight, &res.Available); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (u *Database) GetRooms(ctx context.Context, req *models.GetRoomRequest) (*models.GetRoomResponse, error) {
	query, args, err := sqlbuilder.GetsRoom(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := u.Db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var rooms []*models.UpdateRoomRequest

	for rows.Next() {
		var all models.UpdateRoomRequest

		if err := rows.Scan(&all.HotelID, &all.ID, &all.RoomType, &all.PricePerNight, &all.Available); err != nil {
			log.Println(err)
			return nil, err
		}
		rooms = append(rooms, &all)
	}
	return &models.GetRoomResponse{Rooms: rooms}, nil
}

func (u *Database) UpdateRooms(ctx context.Context, req *models.UpdateRoomRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.UpdateRoom(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int

	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("room has been updated with this id %v", id)}, nil
}

func (u *Database) DeleteRoom(ctx context.Context, req *models.GetRoomRequest) (*models.GeneralResponse, error) {
	query, args, err := sqlbuilder.DeleteRoom(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Room has been deleted with this id %v", req.ID)}, nil
}
