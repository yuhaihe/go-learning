package variable

import "testing"

func TestZeroValues(t *testing.T) {
	n, s, ok := ZeroValues()
	if n != 0 || s != "" || ok {
		t.Fatalf("零值应为 0、空字符串、false，得到 %d %q %v", n, s, ok)
	}
}
