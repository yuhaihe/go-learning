package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UserHandler 处理 GET /users/:id，返回 JSON {"id":"..."}。
func UserHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"id": c.Param("id")},
	)
}

// NewRouter 注册 GET /users/:id。
func NewRouter() *gin.Engine {
	r := gin.New()
	r.GET("/users/:id", UserHandler)
	r.Run("127.0.0.1:3014")
	return r
}

// GetText 在 timeout 内请求 URL。只接受 200，返回响应正文。
func GetText(ctx context.Context, rawURL string, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}
	return string(body), nil
}
