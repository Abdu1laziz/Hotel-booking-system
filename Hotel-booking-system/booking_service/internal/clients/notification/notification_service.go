package notification17

import (
	notificationss "booking-service/pkg/protos/notification"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() notificationss.NotificationClient {
	conn, err := grpc.NewClient("localhost:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("notification error", err)
	}
	client := notificationss.NewNotificationClient(conn)
	return client
}
