package handler

import "github.com/gin-gonic/gin"

//RouteHandler for handleing routes of some-route
func RouteHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello from some-route",
	})
}
