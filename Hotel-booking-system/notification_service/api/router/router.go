package router

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"notification-service/config"
	"notification-service/internal/connections"
	"notification-service/pkg/proto/notification"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRouter() {
	r := mux.NewRouter()
	a := connections.NewService().W
	r.HandleFunc("/ws", a.HandleWebSocket)
	certfile := "./cert/notif.pem"
	keyfile := "./cert/notif-key.pem"
	go Grpc()

	fmt.Println("server started on port 8083")
	log.Fatal(http.ListenAndServeTLS(":8083", certfile, keyfile, r))
}

func Grpc() {
	c := config.Configuration()
	ls, err := net.Listen(c.User.Host, c.User.Port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	server := connections.NewService()
	notification.RegisterNotificationServer(s, server)
	reflection.Register(s)
	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
