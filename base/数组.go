package main

import "fmt"

func main() {
	// 声明数组 var balance [10]float32
	// var numbers = [5]int{1, 2, 3, 4, 5}
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	// 如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
	balance2 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance2[1])

	// 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
	var array = []int{1, 2, 3, 4, 5}
	fmt.Println(array, "array")

	// 访问数组元素
	var salary int = array[0]
	fmt.Println(salary, "salary")
}
