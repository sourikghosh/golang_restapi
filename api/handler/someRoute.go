package handler

import (
	"restapi/api/controller"

	"github.com/gin-gonic/gin"
)

//RouteHandler for handleing routes of some-route
func RouteHandler(ctx *gin.Context) {
	tokenAuth, err := controller.ExtractTokenMetadata(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID, err := controller.FetchSetTokens(ctx, tokenAuth)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.JSON(200, gin.H{
		"success": userID,
	})
}
