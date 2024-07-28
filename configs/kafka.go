package configs

import (
	"github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer
var KafkaReader *kafka.Reader

func InitKafkaWriter() {
	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "messages",
		Balancer: &kafka.LeastBytes{},
	}
}

func InitKafkaReader() {
	KafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "messages",
		GroupID:  "message-group",
		MinBytes: 1e3,
		MaxBytes: 10e6,
	})
}
