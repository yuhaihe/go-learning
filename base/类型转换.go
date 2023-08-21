package main

// Go 语言类型转换基本格式如下：
// type_name(expression)
import "fmt"
import "strconv"

type Writer interface {
	Write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func main() {
	var sum int = 17
	var count int = 5
	var mean float32
	// 整型转化为浮点型
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	// 将一个字符串转换成另一个类型，可以使用以下语法：
	str := "123"
	num, err := strconv.Atoi(str) //strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	// 整数转换为字符串：
	num2 := 123
	str2 := strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num2, str2)

	// 浮点数转换为字符串
	num3 := 3.14
	str3 := strconv.FormatFloat(num3, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num3, str3)

	// 接口类型转换
	// 接口类型转换有两种情况：类型断言和类型转换。

	// 类型断言用于将接口类型转换为指定类型
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
	// 类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：T(value)

	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "Hello, World"
	fmt.Println(sw.str)
}
