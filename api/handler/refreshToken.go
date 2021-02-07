package handler

import (
	"time"

	"restapi/api/controller"

	"github.com/gin-gonic/gin"
)

//Refresh is post
func Refresh(ctx *gin.Context) {
	tokenInfo := ctx.MustGet("tokenInfo").(*controller.RedisTokenDetails)

	deleted, delErr := controller.DeleteAuth(ctx, tokenInfo)
	if delErr != nil || deleted == 0 { //if anything goes wrong
		ctx.AbortWithStatusJSON(500, "Internel Server Error")
		return
	}

	tokenDetails, err := controller.CreateToken(tokenInfo.UserID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	errToken := controller.SetToken(ctx, tokenInfo.UserID, tokenDetails)
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
	ctx.JSON(200, gin.H{
		"access_token": tokenDetails.AccessToken,
	})
}
