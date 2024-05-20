package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

const (
	broker = "localhost:19093"
	group  = "test-group"
)

var Producer sarama.SyncProducer

func CreateNewProducer() sarama.SyncProducer {
	// Create a new Kafka configuration
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return nil
	}
	fmt.Println("Successfully create new kafka producer")
	Producer = producer
	return producer
}

func GetSyncProducer() sarama.SyncProducer {
	return Producer
}
