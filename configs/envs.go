package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser string
	DBPass string
	DBAddr string
	DBName string

	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		// Server
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8000"),

		// DB
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASSWORD", "mypassword"),
		DBAddr: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "vitshop"),

		// Auth
		JWTSecret:              getEnv("JWT_SECRET", "nevergonnagiveyouup"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}
		return num
	}
	return fallback
}
