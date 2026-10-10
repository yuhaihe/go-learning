package flow

import "testing"

func TestSumEven(t *testing.T) {
	if got := SumEven(6); got != 12 {
		t.Fatalf("SumEven(6)=%d，期望 12", got)
	}
	if got := SumEven(-1); got != 0 {
		t.Fatalf("负数应返回 0，得到 %d", got)
	}
}

func TestSign(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{{-2, "neg"}, {0, "zero"}, {3, "pos"}}
	for _, tc := range cases {
		if got := Sign(tc.n); got != tc.want {
			t.Fatalf("Sign(%d)=%s，期望 %s", tc.n, got, tc.want)
		}
	}
}
