package middleware

import (
	"restapi/api/controller"
	"restapi/database"
	"restapi/models"
	"strings"

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
		_, exits := database.GetByEmail(ctx, signupData.Email)
		if exits {
			ctx.AbortWithStatusJSON(400, gin.H{"err": "Email already exists"})
			return
		}
		ctx.Set("data", signupData)
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
		ok := controller.CompareHash(data.Password, loginData.Password)
		if !ok {
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
		accessToken, errAccess := controller.VerifyAccessToken(ctx)
		if errAccess != nil {
			ctx.AbortWithStatusJSON(401, errAccess.Error())
			return
		}
		refreshToken, errRefresh := controller.VerifyRefreshToken(ctx)
		if errRefresh != nil {
			ctx.AbortWithStatusJSON(401, errRefresh.Error())
			return
		}
		accesstokenInfo, errExtraction := controller.ExtractAccessTokenMetadata(accessToken)
		if errExtraction != nil {
			ctx.AbortWithStatusJSON(500, errExtraction.Error())
			return
		}
		refreshtokenInfo, errEx := controller.ExtractRefreshTokenMetadata(refreshToken)
		if errEx != nil {
			ctx.AbortWithStatusJSON(500, errEx.Error())
			return
		}
		errFromRedisAccess := controller.FetchAccessTokens(ctx, accesstokenInfo)
		if errFromRedisAccess != nil {
			ctx.AbortWithStatusJSON(401, errFromRedisAccess.Error())
			return
		}

		errFromRedis := controller.FetchRefreshTokens(ctx, refreshtokenInfo)
		if errFromRedis != nil {
			ctx.AbortWithStatusJSON(401, errFromRedis.Error())
			return
		}
		tokenInfo := accesstokenInfo
		tokenInfo.RefreshUUID = refreshtokenInfo.RefreshUUID
		ctx.Set("tokenInfo", tokenInfo)
		ctx.Next()
	}
}

//AccessTokenAuth checks the AccesToken if its Valid of not
func AccessTokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := controller.VerifyAccessToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(401, err.Error())
			return
		}
		tokenInfo, errExtraction := controller.ExtractAccessTokenMetadata(token)
		if errExtraction != nil {
			ctx.AbortWithStatusJSON(500, errExtraction.Error())
			return
		}
		errFromRedis := controller.FetchAccessTokens(ctx, tokenInfo)
		if errFromRedis != nil {
			ctx.AbortWithStatusJSON(401, errFromRedis.Error())
			return
		}
		ctx.Set("userid", tokenInfo.UserID)
		ctx.Next()
	}
}

//RefreshTokenAuth checks
func RefreshTokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, errRefresh := controller.VerifyRefreshToken(ctx)
		if errRefresh != nil {
			ctx.AbortWithStatusJSON(401, errRefresh.Error())
			return
		}
		refreshtokenInfo, errEx := controller.ExtractRefreshTokenMetadata(refreshToken)
		if errEx != nil {
			ctx.AbortWithStatusJSON(500, errEx.Error())
			return
		}
		errFromRedis := controller.FetchRefreshTokens(ctx, refreshtokenInfo)
		if errFromRedis != nil {
			ctx.AbortWithStatusJSON(401, errFromRedis.Error())
			return
		}

		tokenInfo := refreshtokenInfo
		tokenInfo.AccessUUID = strings.Split(tokenInfo.RefreshUUID, "++")[0]
		ctx.Set("tokenInfo", tokenInfo)
		ctx.Next()
	}
}
