package pointer

// Swap 交换两个指针指向的值。任一指针为 nil 时什么都不做。
func Swap(a, b *int) {
	if a == nil || b == nil {
		return
	}

	*a, *b = *b, *a
}
