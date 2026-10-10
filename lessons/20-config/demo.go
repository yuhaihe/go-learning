//go:build ignore

package main

import (
	"log/slog"
	"os"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":3000"
	}
	slog.Info("config", "addr", addr, "dsn_set", os.Getenv("MYSQL_DSN") != "")
}
