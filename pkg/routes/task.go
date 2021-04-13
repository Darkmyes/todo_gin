package routes

import (
	"todo_gin/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetTaskRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/tasks")
	{
		userRoutes.GET("", handlers.ListTasks)
		userRoutes.POST("", handlers.RegisterTask)
	}
}
