package handler

import (
	"restapi/api/controller"
	"restapi/models"
	"time"

	"github.com/gin-gonic/gin"
)

//Signup routes
func Signup(ctx *gin.Context) {
	data := ctx.MustGet("data").(models.Signup)
	errSave := controller.CreateUser(ctx, data)
	if errSave != nil {
		ctx.JSON(500, gin.H{
			"error": errSave.Error(),
		})
		return
	}

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
	rtDuration := tokenDetails.RtExpires
	second := time.Now().Add(time.Second).Unix()
	cookieDuration := int(rtDuration / second)
	ctx.SetCookie("jid", tokenDetails.RefreshToken,
		cookieDuration,
		"/api/ref",
		"localhost",
		false,
		true,
	)
	ctx.JSON(201, gin.H{
		"access_token": tokenDetails.AccessToken,
	})
}
