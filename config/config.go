package config

import "os"

//Config varibles
var Config map[string]string = initConfig()

func initConfig() map[string]string {
	config := make(map[string]string)

	port, exits := os.LookupEnv("PORT")
	if !exits {
		port = "4000"
	}
	redisURL, exits := os.LookupEnv("REDIS_URL")
	if !exits {
		redisURL = "redis://localhost:6379/0"
	}
	databaseURL, exits := os.LookupEnv("DATABASE_URL")
	if !exits {
		databaseURL = "postgresql://sourik:neverdie@localhost:5432/login_auth"
	}
	config["REDIS_URL"] = redisURL
	config["PORT"] = port
	config["DATABASE_URL"] = databaseURL
	return config
}
