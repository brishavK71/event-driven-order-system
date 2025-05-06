package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafkaWriter(broker, topic string) {

	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func PublishOrder(order any) {
	msgBytes, err := json.Marshal(order)
	if err != nil {
		log.Println("Failed to marshal order:", err)
		return
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("order"),
		Value: msgBytes,
	})
	if err != nil {
		log.Println("Failed to publish message:", err)
	}
}
