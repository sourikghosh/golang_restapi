package routes

import (
	database "restapi/Database"
	models "restapi/models"

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
	conn, err := database.Dbclient(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, error := conn.Query(c, "INSERT INTO go_userlist(email, password) VALUES ($1,$2)",
		signupData.Email, signupData.Password)
	if error != nil {
		c.JSON(500, gin.H{"err": error.Error()})
		return
	}
	c.JSON(201, gin.H{"success": "😎"})
}
