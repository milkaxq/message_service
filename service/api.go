package service

import (
	"log"
	"message_service/configs"
	"message_service/models"

	"github.com/google/uuid"
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

	return message.ID, nil
}
