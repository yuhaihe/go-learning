package bind

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() { gin.SetMode(gin.TestMode) }

func TestParseCreateNote(t *testing.T) {
	ok := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"a","content":"b"}`))
	ok.Header.Set("Content-Type", "application/json")
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = ok
	got, err := ParseCreateNote(c)
	if err != nil || got.Title != "a" || got.Content != "b" {
		t.Fatalf("绑定失败：%+v %v", got, err)
	}

	bad := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"content":"only"}`))
	bad.Header.Set("Content-Type", "application/json")
	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	c.Request = bad
	if _, err := ParseCreateNote(c); err == nil {
		t.Fatal("缺少 title 应校验失败")
	}
}
