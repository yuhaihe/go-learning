package parallel

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(nil); got != 0 {
		t.Fatalf("空切片=%d", got)
	}
	if got := Sum([]int{1, 2, 3, 4}); got != 10 {
		t.Fatalf("Sum=%d，期望 10", got)
	}
}
