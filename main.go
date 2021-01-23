package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(contex *gin.Context) {
		contex.JSON(200, gin.H{
			"success": "😏",
		})
	})
	app.Run()
}
