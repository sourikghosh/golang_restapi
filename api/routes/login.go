package routes

import (
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//Login function
func Login(c *gin.Context) {
	var loginData models.Login
	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"success": "Login😏",
	})
}
