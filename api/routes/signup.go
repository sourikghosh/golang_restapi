package routes

import (
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//Signup routes
func Signup(c *gin.Context) {
	var signupData models.Signup
	if err := c.ShouldBind(&signupData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if signupData.Password != signupData.ConfirmPassword {
		c.JSON(400, gin.H{"error": "ConfirmPassword doesnot match the Password"})
		return
	}

	c.JSON(201, gin.H{"status": "you are logged in"})
}
