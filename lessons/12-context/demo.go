//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Background：空白的根通知，表示先一直做下去。
	// WithTimeout：1 毫秒后这份 ctx 自动取消。cancel 用来提前取消，也可以释放内部计时器。
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel() // main 结束前一定会调用 cancel()

	<-ctx.Done() // 停在这里，直到超时。Done() 是一个 channel，取消后才能往下走
	fmt.Println(ctx.Err()) // context deadline exceeded
}
