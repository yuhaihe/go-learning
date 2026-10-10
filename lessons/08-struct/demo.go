//go:build ignore

package main

import "fmt"

type Rect struct {
	W, H int
}

func (r Rect) Area() int { return r.W * r.H }

func main() {
	fmt.Println(Rect{W: 3, H: 4}.Area())
}
