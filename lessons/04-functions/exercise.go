package fn

import "errors"

// Divide 做整数除法。b 为 0 时返回错误，商返回 0。
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("this is error!")
	}

	return a / b, nil
}
