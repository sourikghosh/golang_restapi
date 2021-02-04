package handler

import (
	"restapi/api/controller"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//Login function
func Login(ctx *gin.Context) {
	data := ctx.MustGet("data").(models.Signup)
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

	ctx.JSON(200, gin.H{
		"access_token":  tokenDetails.AccessToken,
		"refresh_token": tokenDetails.RefreshToken,
	})

}
