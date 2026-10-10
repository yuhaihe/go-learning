package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint `json:"uid"`
	jwt.RegisteredClaims
}

// Sign 用 HS256 签发，claims 里放 uid 和过期时间。
func Sign(secret string, userID uint, exp time.Time) (string, error) {
	panic("TODO")
}

// UserID 用同一密钥解析。过期或签名不对时返回错误。
func UserID(secret, raw string) (uint, error) {
	panic("TODO")
}
