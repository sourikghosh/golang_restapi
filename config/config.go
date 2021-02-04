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
	accessSecret, exits := os.LookupEnv("ACCESS_SECRET")
	if !exits {
		accessSecret = "5df8faaeaf9507f8782d667dca29ed9aa97157038350838e12633e04aa6bdbad"
	}
	refreshSecret, exits := os.LookupEnv("REFRESH_SECRET")
	if !exits {
		refreshSecret = "1485749ecf7ad282805e5e54e5a5abd6e915541c45af5f8d5047dac2738b3f11"
	}
	config["ACCESS_SECRET"] = accessSecret
	config["REFRESH_SECRET"] = refreshSecret
	config["REDIS_URL"] = redisURL
	config["PORT"] = port
	config["DATABASE_URL"] = databaseURL
	return config
}
