//go:build ignore

package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println("原始:", nums, "len", len(nums), "cap", cap(nums))

	nums = append(nums, 3)
	fmt.Println("append 后:", nums, "len", len(nums), "cap", cap(nums))

	fmt.Println("--- range 遍历 ---")
	for i, v := range nums {
		fmt.Println("下标", i, "值", v)
	}

	fmt.Println("--- 下标 for ---")
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
