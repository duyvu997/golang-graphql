package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"graphql-golang/internal/app/entity"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func CreateUser(ctx context.Context, input entity.CreateUserRequest) (entity.CreateUserResponse, error) {
	broker := "localhost:9092"
	topicToPublish := "test1"
	topicToSubscribe := "test1"

	// Create a producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}

	deliveryChan := make(chan kafka.Event)
	// Produce a message to the topic

	value, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicToPublish, Partition: 0},
		Value:          value,
	}, deliveryChan)
	if err != nil {
		panic(err)
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v", m.TopicPartition.Error)
	}
	close(deliveryChan)
	producer.Close()

	// =================================================================
	config := &kafka.ConfigMap{
		"bootstrap.servers":                     "localhost:9092",
		"group.id":                              "test-group",
		"go.application.rebalance.enable":       true,
		"max.in.flight.requests.per.connection": 1,
	}

	// Create consumer
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{topicToSubscribe}, nil)

	for {

		fmt.Println("Commited", consumer)
		msg, err := consumer.ReadMessage(-1)
		fmt.Println("Commited", msg, err)
		if err != nil {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		com, err := consumer.CommitMessage(msg)
		fmt.Println("Commited", com, err)
		if err != nil {
			fmt.Println("Error", err)
		}

		var data entity.CreateUserResponse

		err = json.Unmarshal(msg.Value, &data)
		if err != nil {
			panic(err)
		}
		data.CreatedAt = time.Now()
		data.UID = 123
		fmt.Println("data", data)

		return data, nil
	}
}
