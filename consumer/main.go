package main

import (
	"context"
	"fmt"
	"log"
	"message_service/configs"
	"message_service/models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	configs.InitDB()
	configs.InitKafkaReader()
}

func main() {
	for {
		m, err := configs.KafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err.Error())
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		if err := configs.DB.Model(&models.Message{ID: string(m.Key)}).Update("processed", true).Error; err != nil {
			log.Printf("Error updating message: %v", err.Error())
		}
	}
}
