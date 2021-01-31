package middleware

import (
	"restapi/database"
	models "restapi/models"

	gin "github.com/gin-gonic/gin"
)

//SignupAuth checks the incomming request and Validates
func SignupAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var signupData models.Signup
		if err := ctx.ShouldBind(&signupData); err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if signupData.Password != signupData.ConfirmPassword {
			ctx.AbortWithStatusJSON(402, gin.H{"error": "ConfirmPassword doesnot match the Password"})
			return
		}
		_, exits := database.GetByEmail(ctx, signupData.Email)
		if exits {
			ctx.AbortWithStatusJSON(400, gin.H{"err": "Email already exists"})
			return
		}
		ctx.Set("data", signupData)
		ctx.Next()
	}
}

//LoginAuth checks the incomming request and Validates
func LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginData models.Login
		if err := ctx.ShouldBind(&loginData); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		data, exits := database.GetByEmail(ctx, loginData.Email)
		if !exits {
			ctx.AbortWithStatusJSON(400, gin.H{"err": "invalid email/password"})
			return
		}
		if data.Password != loginData.Password {
			ctx.AbortWithStatusJSON(400, gin.H{"err": "invalid email/password"})
			return
		}
		ctx.Set("data", loginData)
		ctx.Next()
	}
}
