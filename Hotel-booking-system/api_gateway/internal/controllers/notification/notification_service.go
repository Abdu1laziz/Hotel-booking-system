package notification17

import (
	notifications "api-gateway/pkg/protos/notification"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() notifications.NotificationClient {
	conn, err := grpc.NewClient("localhost:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("notification error", err)
	}
	client := notifications.NewNotificationClient(conn)
	return client
}
