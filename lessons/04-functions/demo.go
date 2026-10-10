//go:build ignore

package main

import (
	"errors"
	"fmt"
)

func main() {
	q, err := divide(10, 2)
	fmt.Println(q, err)
	_, err = divide(1, 0)
	fmt.Println(err)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}
