package main

import (
	"fmt"
	"log"
	"net"
	"user-service/config"
	"user-service/internal/connections"
	"user-service/pkg/proto/user"

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
	user.RegisterUserServer(s, server)
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
