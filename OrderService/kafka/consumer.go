package kafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"OrderService/model"
	"OrderService/service"
)

func ConsumeOrders(topic string, kafkaConfig *kafka.ConfigMap, orderService *service.OrderService) {
	consumer, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()


	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var order model.Order
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				log.Printf("Failed to unmarshal order: %v", err)
				continue
			}

			// Process the order
			response := orderService.

			if response. == -1 {
				log.Printf("Order processing failed: %v", response.Msg)
			} else {
				log.Printf("Order processed successfully: %v", response.Msg)
			}
		} else {
			log.Printf("Consumer error: %v", err)
		}
	}
}
