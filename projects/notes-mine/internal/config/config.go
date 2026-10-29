package config

import "os"

type Config struct {
	HTTPAddr  string
	MySQLDSN  string
	RedisAddr string
	JWTSecret string
}

func Load() Config {
	return Config{
		HTTPAddr:  getenv("HTTP_ADDR", ":3000"),
		MySQLDSN:  os.Getenv("MYSQL_DSN"),
		RedisAddr: getenv("REDIS_ADDR", "127.0.0.1:6379"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}

	return value
}
