//go:build ignore

package main

import "fmt"

func main() {
	// make(chan string, 1)：string 类型的 channel，缓冲区能放 1 个值。
	// 缓冲区还空着时，下面的 ch <- 不会卡住，值先停在缓冲区里，不需要立刻有人接收。
	ch := make(chan string, 1)
	ch <- "ping"

	// select 在多个 case 里挑一个「现在就能做」的。
	// 缓冲区里已经有 "ping"，<-ch 取得到，所以走第一个 case，打印 ping。
	// default 只有其它 case 都做不了时才走。这次缓冲区不空，不会进 default。
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("empty")
	}

	// 对比：同样是容量 1，但还没往里面发送，缓冲区是空的。
	// <-empty 现在取不到值。没有 default 的话，程序会停在这里一直等。
	// 有 default 就立刻走 default，打印 empty，不会卡住。
	empty := make(chan string, 1)
	select {
	case msg := <-empty:
		fmt.Println(msg)
	default:
		fmt.Println("empty")
	}
}
