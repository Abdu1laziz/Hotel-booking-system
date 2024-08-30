package main

import (
	"fmt"
	"hotel-service/config"
	"hotel-service/internal/connections"
	"hotel-service/pkg/proto/hotel"
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
	hotel.RegisterHotelServer(s,server)
	reflection.Register(s)
	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
