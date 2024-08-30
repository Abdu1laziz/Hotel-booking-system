package booking

import (
	"log"
	"notification-service/pkg/proto/booking"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() booking.BookHotelClient {
	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := booking.NewBookHotelClient(conn)
	return client
}
