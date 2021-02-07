package controller

import (
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//CreateUser func
func CreateUser(ctx *gin.Context, userInfo models.Signup) error {
	hpass, errHash := Hash(userInfo.Password)
	if errHash != nil {
		return errHash
	}
	err := database.CreateUser(ctx, userInfo.Email, hpass)
	if err != nil {
		return err
	}
	return nil
}
