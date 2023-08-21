package main

import (
	"fmt"
)

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world")
// 	say("hello")
// }

func sum1(s []int, c chan int) {
	q1 := 0
	for _, v := range s {
		q1 += v
	}
	c <- q1 // 把 q1 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum1(s[:len(s)/2], c)
	go sum1(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
