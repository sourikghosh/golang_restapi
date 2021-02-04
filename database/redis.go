package database

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisConn *redis.Client

//RedisClient creates a new client for redis
func RedisClient(connString string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opt, err := redis.ParseURL(connString)
	if err != nil {
		return err
	}
	redisConn = redis.NewClient(opt)
	errConn := redisConn.Ping(ctx).Err()
	if errConn != nil {
		return errConn
	}
	return nil
}

//Set sets the key with the value of exp
func Set(ctx *gin.Context, key string, value string, exp time.Duration) error {
	err := redisConn.Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

//Get gets the key from the redis store
func Get(ctx *gin.Context, key string) (string, error) {
	value, err := redisConn.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

//Del deletes the key provided from redis database
func Del(ctx *gin.Context, keys ...string) (int64, error) {
	value, err := redisConn.Del(ctx, keys...).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}
