package parsex

import (
	"errors"
	"strconv"
	"testing"
)

func TestParsePositive(t *testing.T) {
	n, err := ParsePositive("12")
	if err != nil || n != 12 {
		t.Fatalf("解析 12 得到 %d %v", n, err)
	}
	_, err = ParsePositive("0")
	if !errors.Is(err, ErrNotPositive) {
		t.Fatalf("0 应是 ErrNotPositive，得到 %v", err)
	}
	_, err = ParsePositive("x")
	if !errors.As(err, new(*strconv.NumError)) {
		t.Fatalf("非法数字应包住 strconv 错误，得到 %v", err)
	}
}
