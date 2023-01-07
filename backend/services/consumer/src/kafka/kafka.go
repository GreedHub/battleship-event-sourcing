package kafkaconnector

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConnector struct {
	producer *kafka.Producer
}

func (k *KafkaConnector) CreateProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "host1:9092,host2:9092",
		"client.id":         socket.gethostname(),
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	k.producer = p
}

func (k *KafkaConnector) ProduceShipEvent(e string) {
	topic := "battleship/ship"
	delivery_chan := make(chan kafka.Event, 10000)
	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny},
		Value: []byte(e)},
		delivery_chan,
	)

	if err != nil {
		fmt.Printf("Failed to send event: %s\n", err)
		os.Exit(1)
	}
}
