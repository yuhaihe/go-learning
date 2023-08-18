package main //定义了包名,每个 Go 应用程序都包含一个名为 main 的包
import "fmt" // 告诉 Go 编译器这个程序需要使用 fmt 包

// main() 是程序开始执行的函数,只能存在一个
func main() {
    fmt.Printf("hello, world\n")
}

// func main() {
//     fmt.Printf("hello, world22")
// }

// init()函数会在main之前执行，可定义多个
func init() {
	fmt.Println("init")
	test()
}

func init() {
	fmt.Println("init2")
}

func test() {
	fmt.Println("test")
}