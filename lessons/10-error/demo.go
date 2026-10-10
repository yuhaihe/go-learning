//go:build ignore

package main

import (
	"errors"
	"fmt"
)

var errDuplicate = errors.New("duplicate")

func main() {
	err := fmt.Errorf("save user: %w", errDuplicate)
	fmt.Println(err)
	// 必须和包进去的是同一个错误值，重新 errors.New 一次对不上
	fmt.Println(errors.Is(err, errDuplicate))

	// e := errors.New("msg")
	// e2 := fmt.Errorf("ctx: %w", e)
	// e3 := fmt.Errorf("only: %s", "e3")
	// e4 := errors.Join(e, e2, e3)
	// fmt.Println(e)
	// fmt.Println(e2)
	// fmt.Println(e3)
	// fmt.Println(e4)
}
