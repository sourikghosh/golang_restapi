package middleware

import (
	"restapi/api/controller"
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//SignupAuthentication checks the incomming request and Validates
func SignupAuthentication() gin.HandlerFunc {

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
		data, exits := database.GetByEmail(ctx, signupData.Email)
		if exits {
			ctx.AbortWithStatusJSON(400, gin.H{"err": "Email already exists"})
			return
		}
		ctx.Set("data", data)
		ctx.Next()
	}
}

//LoginAuthentication checks the incomming request and Validates
func LoginAuthentication() gin.HandlerFunc {
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
		ctx.Set("data", data)
		ctx.Next()
	}
}

//TokenAuth checks the token and checks if token is validor not
func TokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := controller.TokenValidator(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(401, err.Error())
			return
		}
		ctx.Next()
	}
}
