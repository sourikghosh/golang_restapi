package main

import (
	"log"

	"restapi/api/handler"
	"restapi/api/middleware"
	"restapi/config"
	"restapi/database"

	"github.com/gin-gonic/gin"
)

func init() {
	errRedis := database.RedisClient(config.Config["REDIS_URL"])
	if errRedis != nil {
		log.Fatalf("⛔ Redis URI is not valid: %v\n", errRedis)
	} else {
		log.Println("REDIS CONNECTED 🥇")
	}

	errDB := database.InitDB(config.Config["DATABASE_URL"])
	if errDB != nil {
		log.Fatalf("⛔ Unable to connect to database: %v\n", errDB)
	} else {
		log.Println("DATABASE CONNECTED 🥇")
	}
}

func main() {

	app := gin.Default()
	api := app.Group("/api")
	{
		api.POST("/login", middleware.LoginAuthentication(), handler.Login)
		api.POST("/signup", middleware.SignupAuthentication(), handler.Signup)
		protected := api.Group("/protected")
		{
			protected.GET("/some-route", handler.RouteHandler)
		}
	}

	app.Run(":" + config.Config["PORT"])
}
