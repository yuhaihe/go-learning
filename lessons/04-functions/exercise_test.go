package fn

import "testing"

func TestDivide(t *testing.T) {
	q, err := Divide(10, 2)
	if err != nil || q != 5 {
		t.Fatalf("Divide(10,2)=%d %v，期望 5 且无错误", q, err)
	}
	if _, err := Divide(1, 0); err == nil {
		t.Fatal("除以 0 应返回错误")
	}
}
