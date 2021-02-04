package handler

import (
	"restapi/api/controller"
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//Signup routes
func Signup(ctx *gin.Context) {
	data := ctx.MustGet("data").(models.Signup)
	database.CreateUser(ctx, data.Email, data.Password)

	tokenDetails, err := controller.CreateToken(data.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	errToken := controller.SetToken(ctx, data.ID, tokenDetails)
	if errToken != nil {
		ctx.JSON(500, gin.H{
			"error": errToken.Error(),
		})
		return
	}
	ctx.SetCookie("jid", tokenDetails.RefreshToken,
		604800,
		"/api/ref",
		"localhost",
		false,
		true,
	)
	ctx.JSON(201, gin.H{
		"access_token": tokenDetails.AccessToken,
	})
}
