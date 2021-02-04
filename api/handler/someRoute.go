package handler

import (
	"github.com/gin-gonic/gin"
)

//RouteHandler for handleing routes of some-route
func RouteHandler(ctx *gin.Context) {
	userid := ctx.MustGet("userid").(string)

	ctx.JSON(200, gin.H{
		"success": userid,
	})
}
