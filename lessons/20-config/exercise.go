package appcfg

import "log/slog"

type Config struct {
	HTTPAddr string
	MySQLDSN string
}

// Load 读取 HTTP_ADDR 和 MYSQL_DSN。地址为空时用 :3000。
func Load() Config {
	panic("TODO")
}

// LogRequest 打一条包含 method、path、status 的 info 日志。
func LogRequest(log *slog.Logger, method, path string, status int) {
	panic("TODO")
}
