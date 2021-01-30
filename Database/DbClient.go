package database

import (
	"fmt"
	"os"
	"restapi/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Dbclient connection
func Dbclient(ctx *gin.Context) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(ctx, config.Config["DATABASE_URL"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close()
	return conn, nil
}
