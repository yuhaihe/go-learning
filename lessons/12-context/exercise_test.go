package ctxt

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Wait(ctx, time.Second); !errors.Is(err, context.Canceled) {
		t.Fatalf("取消后应返回 Canceled，得到 %v", err)
	}
	if err := Wait(context.Background(), time.Millisecond); err != nil {
		t.Fatalf("正常等待应返回 nil，得到 %v", err)
	}
}
