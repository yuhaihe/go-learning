package config

import "testing"

func TestLoad(t *testing.T) {
	t.Setenv("HTTP_ADDR", ":8080")
	t.Setenv("MYSQL_DSN", "dsn")
	t.Setenv("REDIS_ADDR", "10.0.0.1:6379")
	t.Setenv("JWT_SECRET", "secret")

	cfg := Load()
	if cfg.HTTPAddr != ":8080" || cfg.MySQLDSN != "dsn" || cfg.RedisAddr != "10.0.0.1:6379" || cfg.JWTSecret != "secret" {
		t.Fatalf("%+v", cfg)
	}

	t.Setenv("HTTP_ADDR", "")
	t.Setenv("REDIS_ADDR", "")
	t.Setenv("MYSQL_DSN", "")
	t.Setenv("JWT_SECRET", "")

	cfg = Load()
	if cfg.HTTPAddr != ":3000" || cfg.RedisAddr != "127.0.0.1:6379" || cfg.MySQLDSN != "" || cfg.JWTSecret != "" {
		t.Fatalf("%+v", cfg)
	}
}
