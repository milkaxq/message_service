package configs

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var Kafka *kafka.Conn

func InitKafka() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "messages", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	Kafka = conn
}
