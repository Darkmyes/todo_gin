package routes

import (
	"todo_gin/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetTaskRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/tasks")
	{
		userRoutes.GET("", handlers.ListTasks)
		/*userRoutes.GET("/:id", handlers.FindTasks)
		userRoutes.GET("/search/", handlers.SearchTasks)*/

		userRoutes.POST("", handlers.RegisterTask)
		userRoutes.PUT("/:id", handlers.UpdateTask)
		userRoutes.DELETE("/:id", handlers.DeleteTask)
	}
}
