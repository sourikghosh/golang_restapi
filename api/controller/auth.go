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
		return nil, errUUID
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
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.Config["REFRESH_SECRET"]))
	if err != nil {
		return nil, err
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
		return errAccess
	}

	errRefresh := database.Set(ctx, td.RefreshUUID, userid, rt.Sub(now))
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

//ExtractToken extracts the token from the request header "Authorization"
func ExtractToken(ctx *gin.Context) (string, bool) {
	bearerToken := ctx.GetHeader("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1], true
	}
	return "", false
}

//VerifyToken verifies the token from the Header
func VerifyToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString, ok := ExtractToken(ctx)
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

//TokenValidator checks the token if its valid or not and returns error
func TokenValidator(ctx *gin.Context) error {
	token, err := VerifyToken(ctx)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}
	return nil
}

//ExtractTokenMetadata from token
func ExtractTokenMetadata(ctx *gin.Context) (*RedisTokenDetails, error) {
	token, err := VerifyToken(ctx)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}

		return &RedisTokenDetails{
			AccessUUID:  accessUUID,
			RefreshUUID: refreshUUID,
			UserID:      userID,
		}, nil
	}
	return nil, err
}

//FetchSetTokens from database
func FetchSetTokens(ctx *gin.Context, authD *RedisTokenDetails) (string, error) {
	userID, err := database.Get(ctx, authD.AccessUUID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

//DeleteAuth deletes the token uuid provided
func DeleteAuth(ctx *gin.Context, authD *RedisTokenDetails) (int64, error) {
	deleted, err := database.Del(ctx, authD.AccessUUID, authD.RefreshUUID)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
