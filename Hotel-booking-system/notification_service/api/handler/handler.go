package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	email1 "notification-service/utils/email"
	"notification-service/pkg/proto/booking"
	"notification-service/pkg/proto/hotel"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
)

type WebSocket struct {
	Map     map[string]*websocket.Conn
	Mutex   *sync.Mutex
	Hotel   hotel.HotelClient
	Booking booking.BookHotelClient
	Ctx     context.Context
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (u *WebSocket) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket is working")
	userID := r.Header.Get("id")
	fmt.Println(userID)
	if userID == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Current connections map:", u.Map)

	kafkaReader, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("notification"),
		kgo.ConsumerGroup("my_group"),
	)
	if err != nil {
		log.Println("Kafka client creation error:", err)
		return
	}
	defer kafkaReader.Close()
	for {
		fetches := kafkaReader.PollFetches(r.Context())
		if fetches.IsClientClosed() {
			break
		}
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			message := record.Value
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Error writing message to WebSocket:", err)
				return
			}
			u.RoomChecker(conn)
			fmt.Println("room check is over")
		}
	}
}

func (u *WebSocket) AddUser(userID string, conn *websocket.Conn) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, exists := u.Map[userID]; exists {
		log.Printf("User %s is already connected", userID)
		return errors.New("user already exists")
	}
	u.Map[userID] = conn
	log.Printf("User %s added to the map", userID)
	return nil
}


func (u *WebSocket) RoomChecker(conn *websocket.Conn) {
	res, err := u.Hotel.Gets(u.Ctx, &hotel.GetsRequest{})
	if err != nil {
		log.Println(err)
	}
	for _, v := range res.Hotels {
		u.RoomDetective(int(v.Id), conn)
	}

}

func (u *WebSocket) RoomDetective(id int, conn *websocket.Conn) {
	rooms, err := u.Hotel.GetRooms(u.Ctx, &hotel.GetroomRequest{HotelId: int32(id)})
	if err != nil {
		log.Println(err)
	} else {
		for _, v := range rooms.Rooms {
			if v.Available {
				if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Hotel ID %v\nRoom ID %v\nRoom Type %v\nRoom Price Per Night %v\nRoom Available %v\n", v.HotelId, v.Id, v.RoomType, v.PricePerNight, v.Available))); err != nil {
					log.Println("Error writing message to WebSocket:", err)
				}
				u.WaitingUsers(fmt.Sprintf("Hotel ID %v\nRoom ID %v\nRoom Type %v\nRoom Price Per Night %v\nRoom Available %v\n", v.HotelId, v.Id, v.RoomType, v.PricePerNight, v.Available))
			}
		}
	}

}

func (u *WebSocket) WaitingUsers(body string) {
	users, err := u.Booking.Getall(u.Ctx, &booking.Request{})
	if err != nil {
		log.Println(err)
	} else {
		for _, v := range users.Users {
			fmt.Println(v.UserEmail)
			if err := email1.Sent(v.UserEmail, body); err != nil {
				log.Println(err)
			}
		}
	}

}
