package models

type CreateHotelRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Rating   int32  `json:"rating"`
	Address  string `json:"address"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}

type GetsRequest struct {
}

type UpdateHotelRequest struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Rating   int32  `json:"rating"`
	Address  string `json:"address"`
}

type GetHotelRequest struct {
	ID int32 `json:"id"`
}

type GetHotelResponse struct {
	ID       int32                `json:"id"`
	Name     string               `json:"name"`
	Location string               `json:"location"`
	Rating   int32                `json:"rating"`
	Address  string               `json:"address"`
	Rooms    []*UpdateRoomRequest `json:"rooms"`
}

type CreateRoomRequest struct {
	HotelID       int32   `json:"hotel_id"`
	RoomType      string  `json:"room_type"`
	PricePerNight float32 `json:"price_per_night"`
}

type GetRoomRequest struct {
	HotelID   int32 `json:"hotel_id"`
	ID        int32 `json:"id"`
	Available bool  `json:"available"`
}

type GetRoomResponse struct {
	Rooms []*UpdateRoomRequest `json:"rooms"`
}

type UpdateRoomRequest struct {
	Available     bool    `json:"available"`
	RoomType      string  `json:"room_type"`
	PricePerNight float32 `json:"price_per_night"`
	ID            int32   `json:"id"`
	HotelID       int32   `json:"hotel_id"`
}
