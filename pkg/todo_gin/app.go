package todo_gin

import (
	"todo_gin/pkg/routes"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	apiRouter := router.Group("/api")
	{
		routes.SetTaskRoutes(apiRouter)
		routes.SetUserRoutes(apiRouter)
	}

	router.Run(":3000")
}
