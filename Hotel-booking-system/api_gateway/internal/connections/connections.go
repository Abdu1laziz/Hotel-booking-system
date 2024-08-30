package connections

import (
	"api-gateway/api/handler"
	broad "api-gateway/internal/broadcast"
	"api-gateway/internal/controllers/booking"
	hotels "api-gateway/internal/controllers/hotel"
	users "api-gateway/internal/controllers/user"
	redmet "api-gateway/pkg/redis/method"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// NewBroadcast initializes a new broadcast Adjust instance.
func NewBroadcast() *broad.Adjust {
	u := users.UserClient()
	h := hotels.Hotel()
	b := booking.Hotel()
	r := NewRedis()
	ctx := context.Background()
	return &broad.Adjust{U: u, Ctx: ctx, R: r, H: h, B: b}
}

// NewHandler initializes a new handler instance.
func NewHandler() *handler.Handler {
	broadcast := NewBroadcast()
	return &handler.Handler{B: broadcast}
}

// NewRedis initializes a new Redis client.
func NewRedis() *redmet.Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	if err := pingRedis(client, ctx); err != nil {
		log.Fatal(err)
	}
	return &redmet.Redis{R: client, Ctx: ctx}
}

// pingRedis checks the connection to the Redis server.
func pingRedis(client *redis.Client, ctx context.Context) error {
	_, err := client.Ping(ctx).Result()
	return err
}
