package appcfg

import (
	"bytes"
	"log/slog"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("MYSQL_DSN", "dsn")
	cfg := Load()
	if cfg.HTTPAddr != ":3000" || cfg.MySQLDSN != "dsn" {
		t.Fatalf("%+v", cfg)
	}
	t.Setenv("HTTP_ADDR", ":8080")
	if Load().HTTPAddr != ":8080" {
		t.Fatal("应读到 HTTP_ADDR")
	}
}

func TestLogRequest(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	LogRequest(log, "GET", "/health", 200)
	text := buf.String()
	for _, part := range []string{"GET", "/health", "200"} {
		if !bytes.Contains([]byte(text), []byte(part)) {
			t.Fatalf("日志缺少 %s：%s", part, text)
		}
	}
}
