package main

import "fmt"

/*
指针:个指针变量指向了一个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：
var ip *int        指向整型
var fp *float32    指向浮点型
当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
*/
const MAX int = 3

func main() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)

	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)
	//	空指针判断：
	//	if(ptr != nil)     /* ptr 不是空指针 */
	// if(ptr == nil)    /* ptr 是空指针 */

	a2 := []int{10, 100, 200}
	var i int
	var ptr2 [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr2[i] = &a2[i] /* 整数地址赋值给指针数组 */
		// fmt.Printf("ptr2 的值为 : %d\n", ptr2[i])
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a2[%d] = %d\n", i, *ptr2[i])
	}
}

func init() {
	// 指向指针的指针

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

// 语言指针作为函数参数
func init() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
