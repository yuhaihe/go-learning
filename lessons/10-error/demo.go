//go:build ignore

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("save user: %w", errors.New("duplicate"))
	fmt.Println(err)
	fmt.Println(errors.Is(err, errors.New("duplicate")))
}
