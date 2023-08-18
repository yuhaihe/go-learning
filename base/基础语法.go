package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("菜鸟教程：runoob.com")
	// 字符串连接
	fmt.Println("Google" + "Runoob")
	// 格式化字符串
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var num = 10
	var url = "Code=%d&endDate=%s&num=%d"
	// Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
	// Printf 根据格式化参数生成格式化的字符串并写入标准输出。
	var target_url = fmt.Sprintf(url, stockcode, enddate, num)
	fmt.Println(target_url)

	fmt.Printf(url, stockcode, enddate, num)
}
