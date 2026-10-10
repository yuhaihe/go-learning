//go:build ignore

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 6; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
