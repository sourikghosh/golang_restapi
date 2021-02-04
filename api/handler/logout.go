package handler

import (
	"restapi/api/controller"

	"github.com/gin-gonic/gin"
)

//Logout is delete route
func Logout(ctx *gin.Context) {
	tokenInfo := ctx.MustGet("tokenInfo").(*controller.RedisTokenDetails)
	ctx.SetCookie(
		"jid", "",
		-1,
		"/api/ref",
		"localhost",
		false,
		true,
	)
	deleted, delErr := controller.DeleteAuth(ctx, tokenInfo)
	if delErr != nil || deleted == 0 { //if anything goes wrong
		ctx.AbortWithStatusJSON(500, "Internel Server Error")
		return
	}
	ctx.JSON(204, deleted)

}
