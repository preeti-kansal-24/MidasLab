package utils

import (
	"github.com/IBM/sarama"
	"preeti-kansal-24/MidasLab.git/kafka"
	"time"
)

func PublishMessage(key string, message []byte, topic string) {

	keyStrEncoder := sarama.StringEncoder(key)
	messageByteEncoder := sarama.ByteEncoder(message)

	kafka.GetSyncProducer().SendMessage(&sarama.ProducerMessage{
		Topic:     topic,
		Key:       keyStrEncoder,
		Value:     messageByteEncoder,
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Now(),
	})
}
