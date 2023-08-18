package main

import "fmt"

func main() {
	// 变量
	// var a string = "Runoob"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b, c)

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	// b = 13332
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	// c = true
	fmt.Println(c)

	// var intVal int = 2
	// 等同于上面  := 是一个声明语句
	intVal := 2
	fmt.Println(intVal)
}

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

// 这种不带声明格式的只能在函数体中出现
// g, h := 123, "hello"
var v = 100 //函数外变量可以不被使用

func init() {
	g, h := 123, "hello"
	// var v = 100 函数内变量必须被使用
	fmt.Println(x, y, a, b, c, d, e, f, g, h, r)
}

var r = 100

// 值类型和引用类型
// 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：
// 当使用等号 = 将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝：

// 更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
// 一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。
// 这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。
// 当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。
