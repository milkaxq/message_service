package main

import (
	"log"
	"message_service/configs"
	"message_service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	configs.InitDB()
	configs.Migrate(configs.DB)
	configs.InitKafkaWriter()
}

func main() {
	r := gin.Default()

	routes.ApiRoutes(r)

	r.Run()
}
