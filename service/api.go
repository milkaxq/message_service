package service

import (
	"context"
	"log"
	"message_service/configs"
	"message_service/models"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func CreateMessage(content string) (string, error) {
	var message = models.Message{
		ID:        uuid.New().String(),
		Content:   content,
		Processed: false,
	}

	if err := configs.DB.Create(&message).Error; err != nil {
		log.Printf("Error in creating message in database %v", err)
		return "", err
	}

	err := configs.KafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(message.ID),
			Value: []byte(message.Content),
		},
	)

	if err != nil {
		log.Printf("Error in sending message to kafka %v", err)
		return "", err
	}

	return message.ID, nil
}

func GetStats() (int64, error) {
	var count int64

	if err := configs.DB.Model(&models.Message{}).Where("processed = ?", "true").Count(&count).Error; err != nil {
		log.Printf("Error getting stats of processed messages: %v", err)
		return 0, err
	}

	return count, nil
}
