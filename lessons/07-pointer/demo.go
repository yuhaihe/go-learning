//go:build ignore

package main

import "fmt"

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
