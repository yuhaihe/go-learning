package main

import "fmt"

/*
不同类型的局部和全局变量默认值为：
int	0
float32	0
pointer	nil
*/
const Tal float32 = 17610270818 // 全局变量

func main() {
	var name = "hayho" // 局部变量
	fmt.Println(name)
	fmt.Println(sum(1, 2))
}

/* 函数定义-两数相加 */
func sum(a, b int) int { // 形式参数
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}
