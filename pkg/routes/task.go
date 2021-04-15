package routes

import (
	"todo_gin/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetTaskRoutes(router *gin.RouterGroup) {
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.GET("", handlers.ListTasks)
		/*taskRoutes.GET("/:id", handlers.FindTasks)
		taskRoutes.GET("/search/", handlers.SearchTasks)*/

		taskRoutes.POST("", handlers.RegisterTask)
		taskRoutes.PUT("/:id", handlers.UpdateTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}
}
