package routes

import "github.com/gin-gonic/gin"

//Signup routes
func Signup(c *gin.Context) {
	c.JSON(201, gin.H{
		"success": "Signup😏",
	})
}
