package handler

import (
	"restapi/api/controller"

	"github.com/gin-gonic/gin"
)

//Logout is delete route
func Logout(ctx *gin.Context) {

	au, err := controller.ExtractTokenMetadata(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(401, "unauthorized")
		return
	}
	deleted, delErr := controller.DeleteAuth(ctx, au.AccessUUID)
	if delErr != nil || deleted == 0 { //if anything goes wrong
		ctx.AbortWithStatusJSON(401, "unauthorized")
		return
	}
	ctx.JSON(200, deleted)

}
