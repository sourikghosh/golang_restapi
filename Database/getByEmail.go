package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

//GetByEmail sends if Email exits
func GetByEmail(c *gin.Context, client *pgxpool.Pool, email string) bool {
	//rows, err := client.Query(c, "Select id from go_userlist where email = $1", email)
	return true
}
