package controller

import (
	"fmt"
	"strings"
	"time"

	"restapi/config"
	"restapi/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

//RedisTokenDetails struct
type RedisTokenDetails struct {
	AccessUUID  string
	RefreshUUID string
	UserID      string
}

//TokenDetails struct
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

//CreateToken will generate the accesstoken and refreshtoken with their respective claims
//and return *TokenDetails , Error
func CreateToken(userid string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	AccessUUID, errUUID := uuid.NewV4()
	if errUUID != nil {
		return nil, errors.Wrap(errUUID, "Could not create uuid for accessToken")
	}
	td.AccessUUID = AccessUUID.String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = td.AccessUUID + "++" + userid

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.Config["ACCESS_SECRET"]))
	if err != nil {
		return nil, errors.Wrap(err, "AccessToken signing failed")
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.Config["REFRESH_SECRET"]))
	if err != nil {
		return nil, errors.Wrap(err, "RefreshToken signing failed")
	}
	return td, nil
}

//SetToken to redis database
func SetToken(ctx *gin.Context, userid string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()
	errAccess := database.Set(ctx, td.AccessUUID, userid, at.Sub(now))
	if errAccess != nil {
		return errors.Wrap(errAccess, "AccessToken signing failed")
	}

	errRefresh := database.Set(ctx, td.RefreshUUID, userid, rt.Sub(now))
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

//ExtractAccessToken extracts the token from the request header "Authorization"
func ExtractAccessToken(ctx *gin.Context) (string, bool) {
	bearerToken := ctx.GetHeader("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1], true
	}
	return "", false
}

//VerifyAccessToken verifies the token from the Header
func VerifyAccessToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString, ok := ExtractAccessToken(ctx)
	if !ok {
		return nil, fmt.Errorf("Empty Token found")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config["ACCESS_SECRET"]), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//VerifyRefreshToken checks refreshToken
func VerifyRefreshToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString, err := ctx.Cookie("jid")
	if err != nil {
		return nil, fmt.Errorf("Cookie Not Found")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config["REFRESH_SECRET"]), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//ExtractAccessTokenMetadata from token
func ExtractAccessTokenMetadata(token *jwt.Token) (*RedisTokenDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, errors.New("No access_uuid claim found")
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, errors.New("No user_id claim found")
		}

		return &RedisTokenDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, errors.New("Token extraction error")
}

//ExtractRefreshTokenMetadata extracts
func ExtractRefreshTokenMetadata(token *jwt.Token) (*RedisTokenDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, errors.New("No refresh_uuid claim found")
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, errors.New("No user_id claim found")
		}

		return &RedisTokenDetails{
			RefreshUUID: refreshUUID,
			UserID:      userID,
		}, nil
	}
	return nil, errors.New("Token extraction error")
}

//FetchAccessTokens from database
func FetchAccessTokens(ctx *gin.Context, authD *RedisTokenDetails) error {
	userID, err := database.Get(ctx, authD.AccessUUID)
	if err != nil || userID != authD.UserID {
		return err
	}
	return nil
}

//FetchRefreshTokens fetchs
func FetchRefreshTokens(ctx *gin.Context, authD *RedisTokenDetails) error {
	userID, err := database.Get(ctx, authD.RefreshUUID)
	if err != nil || userID != authD.UserID {
		return err
	}
	return nil
}

//DeleteAuth deletes the token uuid provided
func DeleteAuth(ctx *gin.Context, authD *RedisTokenDetails) (int64, error) {
	deleted, err := database.Del(ctx, authD.AccessUUID, authD.RefreshUUID)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
