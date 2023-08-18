package main

import "fmt"
import "unsafe"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	println(a, b, c)
}

// 常量还可以用作枚举：
// const (
// 	Unknown = 0
// 	Female  = 1
// 	Male    = 2
// )

func init() {
	const (
		a = "ab"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	println(a, b, c)
}

// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
func init() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
