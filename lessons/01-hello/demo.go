//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println(hello("Go"))
	fmt.Println(hello(""))
}

func hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return "Hello, " + name
}
