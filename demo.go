package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("msg")
	e2 := fmt.Errorf("ctx: %w", e)
	e3 := fmt.Errorf("only: %s", "e3")
	e4 := errors.Join(e, e2, e3)
	fmt.Println(e)
	fmt.Println(e2)
	fmt.Println(e3)
	fmt.Println(e4)
}
