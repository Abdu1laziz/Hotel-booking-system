package sqlbuilder

import (
	"fmt"
	"hotel-service/models"
	"log"

	"github.com/Masterminds/squirrel"
)

func CreateHotel(req *models.CreateHotelRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("hotels").
		Columns("name", "location", "rating", "address").
		Values(req.Name, req.Location, req.Rating, req.Address).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetHotel(req *models.GetHotelRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("hotels").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetsHotel(req *models.GetsRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("hotels").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func UpdateHotel(req *models.UpdateHotelRequest) (string, []interface{}, error) {
	setMap := make(map[string]interface{})

	if req.Name != "" {
		setMap["name"] = req.Name
	}
	if req.Location != "" {
		setMap["location"] = req.Location
	}
	if req.Rating != 0 {
		setMap["rating"] = req.Rating
	}
	if req.Address != "" {
		setMap["address"] = req.Address
	}

	if len(setMap) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}
	query, args, err := squirrel.Update("hotels").
		SetMap(setMap).
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func DeleteHotel(req *models.GetHotelRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("hotels").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func CreateRoom(req *models.CreateRoomRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("rooms").
		Columns("hotel_id", "room_type", "price_per_night").
		Values(req.HotelID, req.RoomType, req.PricePerNight).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetRoom(req *models.GetRoomRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("rooms").
		Where(squirrel.Eq{"id": req.ID, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetsRoom(req *models.GetRoomRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("rooms").
		Where(squirrel.Eq{"available": true, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}
func UpdateRoom(req *models.UpdateRoomRequest) (string, []interface{}, error) {
	setMap := make(map[string]interface{})

	if req.Available{
		setMap["available"] = req.Available
	}else if !req.Available{
		setMap["available"] = req.Available
	}
	if req.RoomType != ""{
		setMap["room_type"] = req.RoomType
	}
	if req.PricePerNight != 0 {
		setMap["price_per_night"] = req.PricePerNight
	}

	if len(setMap) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	query, args, err := squirrel.Update("rooms").
		SetMap(setMap).
		Where(squirrel.Eq{"id": req.ID, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func DeleteRoom(req *models.GetRoomRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("rooms").
		Where(squirrel.Eq{"id": req.ID, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetRoomForHotel(req int) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("rooms").
		Where(squirrel.Eq{"hotel_id": req}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}
