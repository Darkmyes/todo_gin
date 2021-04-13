package routes

import "github.com/gin-gonic/gin"

func SetUserRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello",
			})
		})
	}
}
