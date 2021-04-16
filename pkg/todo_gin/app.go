package todo_gin

import (
	"todo_gin/pkg/handlers"
	"todo_gin/pkg/routes"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	authMiddleware := routes.GetAuthMiddleware()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	auth := router.Group("/auth")
	auth.POST("/auth/login", authMiddleware.LoginHandler)
	auth.POST("/auth/logout", authMiddleware.LogoutHandler)
	auth.POST("/auth/register", handlers.RegisterUser)
	auth.GET("/auth/refresh_token", authMiddleware.RefreshHandler)

	apiRouter := router.Group("/api")
	apiRouter.Use(authMiddleware.MiddlewareFunc())
	{
		routes.SetTaskRoutes(apiRouter)
		routes.SetUserRoutes(apiRouter)
	}

	router.Run(":3000")
}
