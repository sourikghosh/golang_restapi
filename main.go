package main

import (
	"log"
	"restapi/api/handler"
	"restapi/api/middleware"
	"restapi/config"
	"restapi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("⛔ Unable to connect to database: %v\n", err)
	} else {
		log.Println("DATABASE CONNECTED 🥇")
	}
	app := gin.Default()
	api := app.Group("/api")
	{
		api.POST("/login", middleware.LoginAuth(), handler.Login)
		api.POST("/signup", middleware.SignupAuth(), handler.Signup)
	}

	app.Run(":" + config.Config["PORT"])
}
