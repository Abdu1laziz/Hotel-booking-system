package sqlbuilder

import (
	"booking-service/models"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Create(req *models.BookHotelRequest, roomPrice float64) (string, []interface{}, error) {
	// if req.CheckInDate.Before(time.Now().Truncate(24 * time.Hour)) {
	// 	return "", nil, errors.New("check-in date cannot be in the past")
	// }

	totalCost := TotalCostCalculate(req.CheckInDate, req.CheckOutDate, roomPrice)
	fmt.Println(roomPrice)
	fmt.Println(totalCost)
	fmt.Println(req.CheckInDate)
	fmt.Println(req.CheckOutDate)
	var status string
	today := time.Now().Truncate(24 * time.Hour)
	if req.CheckInDate.Equal(today) {
		status = "Guest Living"
	} else {
		status = "Booked"
	}
	query, args, err := squirrel.Insert("booked").
		Columns("user_id", "hotel_id", "room_id", "room_type", "enterydate", "leavingdate", "totalcost", "status").
		Values(req.UserID, req.HotelID, req.RoomID, req.RoomType, req.CheckInDate, req.CheckOutDate, totalCost, status).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return query, args, nil
}

func Get(req *models.GetUsersBookRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("booked").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func timestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

func Update(req *models.BookHotelUpdateRequest, roomPrice float64) (string, []interface{}, error) {
	setMap := make(map[string]interface{})

	// Convert timestamppb.Timestamp to time.Time
	checkInDate := timestampToTime(req.CheckInDate)
	checkOutDate := timestampToTime(req.CheckOutDate)

	if checkInDate != (time.Time{}) {
		setMap["enterydate"] = checkInDate
	}
	if checkOutDate != (time.Time{}) {
		setMap["leavingdate"] = checkOutDate
	}
	if req.RoomID != 0 {
		setMap["room_id"] = req.RoomID
	}
	if req.RoomType != "" {
		setMap["room_type"] = req.RoomType
	}
	if checkInDate != (time.Time{}) && checkOutDate != (time.Time{}) {
		totalCost := TotalCostCalculate(checkInDate, checkOutDate, roomPrice)
		setMap["totalcost"] = totalCost
	}
	setMap["status"] = "updated"

	// Build SQL query using squirrel
	query, args, err := squirrel.Update("booked").
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

// TotalCostCalculate is a placeholder function; replace it with your actual implementation
//
//	func TotalCostCalculate(checkInDate, checkOutDate time.Time, roomPrice float64) float64 {
//		// Example implementation for calculating total cost
//		duration := checkOutDate.Sub(checkInDate).Hours() / 24
//		return duration * roomPrice
//	}
func Cancel(req *models.CancelRoomRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("booked").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}
func TotalCostCalculate(in, out time.Time, price float64) float64 {
	// Ensure 'out' is after 'in'
	if !out.After(in) {
		fmt.Println("Check-out date must be after check-in date.")
		return 0
	}

	// Calculate the total hours between the two times
	hours := out.Sub(in).Hours()

	// Convert hours to days, rounding up to count partial days as full days
	days := math.Ceil(hours / 24)

	// Calculate total cost
	totalCost := days * price

	fmt.Printf("Check-in: %v, Check-out: %v, Days: %v, Total cost: %v\n", in, out, days, totalCost)

	return totalCost
}

func CreateWaitingList(req *models.CreateWaitingList) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("waitinglist").
		Columns("user_id", "hotel_id", "room_type", "user_email", "enterydate", "leavingdate", "status").
		Values(req.UserID, req.HotelID, req.RoomType, req.UserEmail, req.CheckInDate, req.CheckOutDate, "waiting").
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetWaitingList(req *models.GetWaitinglistRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("waitinglist").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func UpdateWaitingList(req *models.UpdateWaitingListRequest) (string, []interface{}, error) {
	setMap := make(map[string]interface{})
	if req.CheckInDate != (time.Time{}) {
		setMap["enterydate"] = req.CheckInDate
	}
	if req.CheckOutDate != (time.Time{}) {
		setMap["leavingdate"] = req.CheckOutDate
	}
	if req.HotelID != 0 {
		setMap["room_id"] = req.HotelID
	}
	if req.RoomType != "" {
		setMap["room_type"] = req.RoomType
	}
	if req.UserID != 0 {
		setMap["user_id"] = req.UserID
	}
	setMap["status"] = "updated"
	query, args, err := squirrel.Update("waitinglist").
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

func DeleteWaitingList(req *models.DeleteWaitingList) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("waitinglist").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetRoomInfo(req *models.GetRoomInfo) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("booked").
		Where(squirrel.Eq{"hotel_id": req.HotelID, "room_id": req.RoomID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetAllUsers() (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("waitinglist").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}
