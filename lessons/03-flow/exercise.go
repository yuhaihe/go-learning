package flow

// SumEven 把 0 到 n（含）的偶数相加。n < 0 时返回 0。
func SumEven(n int) int {
	if n < 0 {
		return 0
	}
	sum := 0
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}

// Sign 负数返回 neg，0 返回 zero，正数返回 pos。
func Sign(n int) string {
	switch {
	case n < 0:
		return "neg"
	case n == 0:
		return "zero"
	default:
		return "pos"
	}
}
