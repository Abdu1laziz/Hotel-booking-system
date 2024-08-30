package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	// Создаем новый TCP слушатель
	ls, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started on port %s\n", ":8083")

	// Запускаем HTTP сервер
	if err := http.Serve(ls, nil); err != nil {
		log.Fatal(err)
	}
}
