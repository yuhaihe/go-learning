package appcfg

import (
	"log/slog"
	"os"
)

type Config struct {
	HTTPAddr string
	MySQLDSN string
}

// Load 读取 HTTP_ADDR 和 MYSQL_DSN。地址为空时用 :3000。
func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":3000"
	}
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "user:pass@tcp(127.0.0.1:3306)/notes"
	}

	return Config{
		HTTPAddr: addr,
		MySQLDSN: dsn,
	}
}

// LogRequest 打一条包含 method、path、status 的 info 日志。
func LogRequest(log *slog.Logger, method, path string, status int) {
	log.Info("request", "method", method, "path", path, "status", status)
}
