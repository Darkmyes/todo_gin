package routes

import (
	"todo_gin/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", handlers.ListUsers)
		//userRoutes.POST("", handlers.RegisterUser)
		//userRoutes.POST("/login", handlers.Login)

		/*userRoutes.GET("/:id", handlers.FindTasks)
		userRoutes.GET("/search/", handlers.SearchTasks)

		userRoutes.POST("", handlers.RegisterTask)
		userRoutes.PUT("/:id", handlers.UpdateTask)
		userRoutes.DELETE("/:id", handlers.DeleteTask)*/
	}
}
