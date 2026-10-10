package parsex

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrNotPositive 表示解析结果不是正数。
var ErrNotPositive = errors.New("not positive")

// ParsePositive 把字符串解析成正整数。
// 不是数字时包住 strconv 的错误；小于等于 0 时包住 ErrNotPositive。
func ParsePositive(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	if n <= 0 {
		return 0, fmt.Errorf("%w", ErrNotPositive)
	}
	return n, nil
}
