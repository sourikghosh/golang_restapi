package handler

import (
	"github.com/gin-gonic/gin"
)

//Login function
func Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"success": "Logged In😏",
	})
}
