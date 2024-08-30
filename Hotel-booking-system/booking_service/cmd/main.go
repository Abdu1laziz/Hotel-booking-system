package main

import (
	"booking-service/config"
	"booking-service/internal/connections"
	"booking-service/pkg/protos/booking"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Configuration()
	ls, err := net.Listen(c.User.Host, c.User.Port)
	
	if err != nil {
		log.Fatal(err)
	}
	
	s := grpc.NewServer()
	server := connections.NewGrpc()
	booking.RegisterBookHotelServer(s, server)
	reflection.Register(s)
	a := connections.NewConsumer()
	
	go func() {
		a.Consumer()
	}()
	fmt.Printf("server started on the port %s", c.User.Port)


	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}

}
