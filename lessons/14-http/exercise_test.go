package httpx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserRoute(t *testing.T) {
	r := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/users/7", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("状态码 %d", rec.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["id"] != "7" {
		t.Fatalf("id=%q", body["id"])
	}
}

func TestGetText(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	}))
	defer srv.Close()

	got, err := GetText(context.Background(), srv.URL, time.Second)
	if err != nil || got != "pong" {
		t.Fatalf("GetText=%q %v", got, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := GetText(ctx, srv.URL, time.Second); err == nil {
		t.Fatal("已取消的 ctx 应返回错误")
	}
}
