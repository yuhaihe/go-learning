//go:build ignore

package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("dev")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": 7,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	fmt.Println(signed)
}
