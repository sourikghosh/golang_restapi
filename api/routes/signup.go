package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Signup routes
func Signup(c *gin.Context) {
	id, idExits := c.GetPostForm("id")
	email, emailExists := c.GetPostForm("email")
	password, passwordExits := c.GetPostForm("password")
	confirmPassword, confirmPasswordExits := c.GetPostForm("confirmPassword")

	if idExits && emailExists && passwordExits && confirmPasswordExits {
		c.JSON(201, gin.H{
			"success":         "Signup😏",
			"id":              id,
			"email":           email,
			"password":        password,
			"confirmPassword": confirmPassword,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": "Error😏",
		})
	}

}
