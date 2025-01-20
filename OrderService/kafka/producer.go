package kafka

import (
	"encoding/json"
	"log"

	"OrderService/model"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func PublishOrder(order model.Order, topic string, kafkaConfig *kafka.ConfigMap) error {
	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
		return err
	}
	defer producer.Close()

	orderBytes, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Failed to marshal order: %v", err)
		return err
	}

	// Publish the message to the Kafka topic
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          orderBytes,
	}, nil)

	if err != nil {
		log.Printf("Failed to publish order: %v", err)
		return err
	}

	log.Printf("Order published to topic %s: %v", topic, string(orderBytes))
	return nil
}
