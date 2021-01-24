package main

import (
	"os"

	"restapi/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	api := app.Group("/api")
	{
		api.POST("/login", routes.Login)
		api.POST("/signup", routes.Signup)
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}
	app.Run(":" + PORT)
}
