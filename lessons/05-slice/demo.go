//go:build ignore

package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3}
	nums = append(nums, 3)
	fmt.Println(nums, len(nums), cap(nums))
}
