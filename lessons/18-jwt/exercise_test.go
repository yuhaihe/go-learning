package token

import (
	"testing"
	"time"
)

func TestSignAndParse(t *testing.T) {
	signed, err := Sign("secret", 7, time.Now().Add(time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	id, err := UserID("secret", signed)
	if err != nil || id != 7 {
		t.Fatalf("UserID=%d %v", id, err)
	}
	if _, err := UserID("other", signed); err == nil {
		t.Fatal("密钥不对时应失败")
	}
	expired, err := Sign("secret", 7, time.Now().Add(-time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := UserID("secret", expired); err == nil {
		t.Fatal("过期 token 应失败")
	}
}
