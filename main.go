package main

import (
	"restapi/api/routes"
	"restapi/config"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	api := app.Group("/api")
	{
		api.POST("/login", routes.Login)
		api.POST("/signup", routes.Signup)
	}

	app.Run(":" + config.Config["PORT"])
}
