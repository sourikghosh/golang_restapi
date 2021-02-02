package database

import (
	"github.com/go-redis/redis/v8"
)

var redisConn *redis.Client

//RedisClient creates a new client for redis
func RedisClient(connString string) error {
	opt, err := redis.ParseURL(connString)
	if err != nil {
		return err
	}
	redisConn = redis.NewClient(opt)
	return nil
}
