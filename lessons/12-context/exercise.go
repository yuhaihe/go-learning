package ctxt

import (
	"context"
	"time"
)

// Wait 在 ctx 取消时返回 ctx.Err()，否则等待 d 后返回 nil。
func Wait(ctx context.Context, d time.Duration) error {
	panic("TODO")
}
