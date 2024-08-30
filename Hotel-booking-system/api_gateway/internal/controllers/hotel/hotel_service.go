package hotelservice

import (
	"api-gateway/pkg/protos/hotel"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() hotel.HotelClient {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := hotel.NewHotelClient(conn)
	return client
}
