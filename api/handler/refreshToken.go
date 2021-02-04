package handler

import "github.com/gin-gonic/gin"

//Refresh is post
func Refresh(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"success": "hello"})
}
