package main

import "fmt"

func main() {
	// sum := 0
	// 基础
	// for i := 0; i <= 10; i++ {
	// 	sum++
	// }
	// 类似while
	// for sum <= 10 {
	// 	sum++
	// }
	// 死循环
	// for {
	// 	sum++ // 无限循环下去
	// 	fmt.Println(sum)
	// }
	// fmt.Println(sum)

	// For-each range 循环
	// strings := []string{"google", "runoob"}
	// fmt.Println(strings)
	// for i, s := range strings {
	// 	fmt.Println(i, s)
	// }

	// numbers := [6]int{1, 2, 3, 5}
	// fmt.Println(numbers)
	// for i, x := range numbers {
	// 	fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	// }
	// for 循环的 range 格式可以省略 key 和 value，如下实例：
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	fmt.Println(map1)
	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
}
