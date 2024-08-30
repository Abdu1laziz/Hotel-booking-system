package producer

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

func Producer(key, message string) error {
	topic := "notification"
	brokerAddress := "localhost:9092"

	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(brokerAddress),
		kgo.WithLogger(kgo.BasicLogger(os.Stdout, kgo.LogLevelWarn, func() string { return time.Now().Format(time.RFC3339) })),
	)
	if err != nil {
		log.Printf("Failed to create Kafka client: %v", err)
		return err
	}
	defer kafkaClient.Close()

	adminClient := kadm.NewClient(kafkaClient)

	ctx := context.Background()
	topics, err := adminClient.ListTopics(ctx)
	if err != nil {
		log.Printf("Failed to list Kafka topics: %v", err)
		return err
	}

	_, exists := topics[topic]
	if !exists {
		_, err = adminClient.CreateTopics(ctx, 1, 1, nil, topic)
		if err != nil {
			log.Printf("Failed to create Kafka topic: %v", err)
			return err
		}
		log.Printf("Topic %s created", topic)
	} else {
		log.Printf("Topic %s already exists", topic)
	}

	w := kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()

	var i int

	for {
		err := w.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte(key),
				Value: []byte(message),
			},
		)
		if err != nil {
			log.Println("Error writing message:", err)
		} else {
			log.Println("Message sent to Kafka")
			return nil
		}
		i++
		time.Sleep(5 * time.Second)
	}
}
