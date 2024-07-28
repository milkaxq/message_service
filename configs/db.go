package configs

import (
	"fmt"
	"log"
	"message_service/models"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(0)

	DB = db

	log.Println("Database connection established successfully")
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Message{}); err != nil {
		log.Fatalf("Error on migration to database %v", err)
	}
}
