package speak

import "testing"

func TestSay(t *testing.T) {
	if got := Say(Dog{}); got != "woof" {
		t.Fatalf("Say=%q，期望 woof", got)
	}
	if got := Say(nil); got != "" {
		t.Fatalf("nil speaker 应返回空字符串，得到 %q", got)
	}
}
