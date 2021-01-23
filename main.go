package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(contex *gin.Context) {
		contex.JSON(200, gin.H{
			"success": "😏",
		})
	})
	os.Setenv("PORT", "4000")
	PORT := os.Getenv("PORT")
	app.Run(":" + PORT)
}
