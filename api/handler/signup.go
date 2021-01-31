package handler

import (
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//Signup routes
func Signup(ctx *gin.Context) {
	data := ctx.MustGet("data").(models.Signup)

	database.CreateUser(ctx, data.Email, data.Password)
	ctx.JSON(201, gin.H{"success": "Created 😋"})
}
