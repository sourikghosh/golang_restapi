package database

import (
	"context"
	"restapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var conn *pgxpool.Pool

//InitDB initialize the database pool and returns error
func InitDB(connString string) error {
	var err error
	conn, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return err
	}
	return nil
}

//GetByEmail sends if Email exits
func GetByEmail(ctx *gin.Context, email string) (models.Login, bool) {
	var details models.Login
	err := conn.QueryRow(ctx, "Select email,password from go_userlist where email = $1", email).Scan(&details.Email, &details.Password)

	if err != nil {
		return details, false
	}
	return details, true
}

//CreateUser will add the user to the database and take the email and password
func CreateUser(ctx *gin.Context, email string, password string) {

	conn.QueryRow(ctx, "INSERT INTO go_userlist(email, password) VALUES ($1,$2)",
		email, password)
}
