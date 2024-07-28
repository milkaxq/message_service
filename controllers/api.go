package controllers

import (
	"message_service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func CreateMessage(c *gin.Context) {
	var json struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageID, err := service.CreateMessage(json.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": messageID})
}
