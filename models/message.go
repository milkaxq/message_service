package models

import (
	"time"
)

type Message struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	Processed bool      `json:"processed"`
	CreatedAt time.Time `json:"created_at"`
}
