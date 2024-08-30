package connections

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"user-service/config"
	ntf "user-service/internal/client/notifications"
	intr "user-service/internal/interface"
	"user-service/internal/interface/service"
	adj "user-service/internal/service/adjust"
	grpcmet "user-service/internal/service/methods"
	"user-service/pkg/database/adjust"
	cons "user-service/pkg/kafka/consumer"
	"user-service/pkg/proto/notification"

	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
	N  notification.NotificationClient
}

// Реализация методов интерфейса intr.User

var _ intr.User = &adjust.Database{}

func NewDatabase() (intr.User, error) {
	c := config.Configuration()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Println("Error pinging database:", err)
		return nil, err
	}
	n := ntf.Hotel()

	// Проверьте, что adjust.Database реализует интерфейс intr.User
	database := &adjust.Database{Db: db, N: n}
	return database, nil
}

func NewService() (*service.Service, error) {
	a, err := NewDatabase()
	if err != nil {
		return nil, err
	}
	return &service.Service{D: a}, nil
}

func NewAdjust() (intr.Adjust, error) {
	a, err := NewService()
	if err != nil {
		return nil, err
	}
	return &adj.Adjust{S: a}, nil
}

func NewAdjustService() (*service.Adjust, error) {
	a, err := NewAdjust()
	if err != nil {
		return nil, err
	}
	return &service.Adjust{A: a}, nil
}

func NewGrpc() (*grpcmet.Service, error) {
	a, err := NewAdjustService()
	if err != nil {
		return nil, err
	}
	return &grpcmet.Service{S: a}, nil
}

func NewConsumer() (*cons.Consumer1, error) {
	a, err := NewGrpc()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return &cons.Consumer1{C: a, Ctx: ctx}, nil
}
