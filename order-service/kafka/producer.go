package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafkaWriter(broker, topic string) {

	// ensureTopicExists(broker, topic)

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

// func ensureTopicExists(broker, topic string) {
// 	conn, err := kafka.Dial("tcp", broker)
// 	if err != nil {
// 		log.Fatalf("❌ Failed to connect to Kafka broker: %v", err)
// 	}
// 	defer conn.Close()

// 	controller, err := conn.Controller()
// 	if err != nil {
// 		log.Fatalf("❌ Failed to get controller: %v", err)
// 	}

// 	controllerConn, err := kafka.Dial("tcp", controller.Host+":"+strconv.Itoa(controller.Port))
// 	if err != nil {
// 		log.Fatalf("❌ Failed to connect to controller: %v", err)
// 	}
// 	defer controllerConn.Close()

// 	topicConfigs := []kafka.TopicConfig{
// 		{
// 			Topic:             topic,
// 			NumPartitions:     1,
// 			ReplicationFactor: 1,
// 		},
// 	}

// 	err = controllerConn.CreateTopics(topicConfigs...)
// 	if err != nil {
// 		log.Printf("⚠️ Could not create topic %s (maybe it already exists): %v", topic, err)
// 	} else {
// 		log.Printf("✅ Topic '%s' created successfully", topic)
// 	}
// }
