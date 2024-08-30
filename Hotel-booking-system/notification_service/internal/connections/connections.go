package connections

import (
	"context"
	"notification-service/api/handler"
	"notification-service/internal/clients/booking"
	"notification-service/internal/clients/hotel"
	"notification-service/internal/services"
	"sync"

	"github.com/gorilla/websocket"
)

func NewWebSocket() *handler.WebSocket{
	h:=hotel.Hotel()
	b:=booking.Hotel()
	ctx:=context.Background()
	return &handler.WebSocket{
		Map:   make(map[string]*websocket.Conn),
		Mutex: &sync.Mutex{},    
		Hotel: h,
		Booking: b,
		Ctx: ctx,           
	}
}

func NewService()*services.Service{
	a:=NewWebSocket()
	return &services.Service{W: a}
}