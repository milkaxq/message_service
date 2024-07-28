package routes

import (
	"message_service/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	api := r.Group("api")
	{
		api.GET("/ping", controllers.PingPong)
		api.POST("/messages", controllers.CreateMessage)
	}
}
