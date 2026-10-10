//go:build ignore

package main

import "fmt"

func main() {
	score := map[string]int{"a": 1}
	score["b"] = 2
	if v, ok := score["c"]; !ok {
		fmt.Println("missing", v)
	}
	fmt.Println(score)
}
