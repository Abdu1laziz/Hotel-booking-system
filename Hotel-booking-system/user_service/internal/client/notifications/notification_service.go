package notification

import (
	"log"
	ntf "user-service/pkg/proto/notification"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() ntf.NotificationClient {
	// Устанавливаем соединение с gRPC сервером
	conn, err := grpc.Dial("localhost:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error connecting to notification service:", err)
		return nil // Возвращаем nil в случае ошибки
	}

	// Создаем новый клиент для уведомлений
	client := ntf.NewNotificationClient(conn)
	return client
}
