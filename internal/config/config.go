package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTSecret string
	JWTTTLMinutes int
}

func Load() *Config {
	_ = godotenv.Load()
	minutes, _ := strconv.Atoi(getenv("JWT_TTL_MINUTES", "120"))
	cfg := &Config{
		AppEnv: getenv("APP_ENV", "development"),
		AppPort: getenv("APP_PORT", "8080"),
		DBHost: getenv("DB_HOST", "localhost"),
		DBPort: getenv("DB_PORT", "3306"),
		DBUser: getenv("DB_USER", "root"),
		DBPass: getenv("DB_PASS", ""),
		DBName: getenv("DB_NAME", "appdb"),
		JWTSecret: getenv("JWT_SECRET", "dev-secret"),
		JWTTTLMinutes: minutes,
	}
	log.Printf("[config] loaded env=%s port=%s db=%s@%s:%s/%s", cfg.AppEnv, cfg.AppPort, cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)
	return cfg
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" { return v }
	return def
}
