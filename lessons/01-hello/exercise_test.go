package hello

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("Go"); got != "Hello, Go" {
		t.Fatalf("Hello(Go)=%q，期望 Hello, Go", got)
	}
	if got := Hello(""); got != "Hello, World" {
		t.Fatalf("空名字应返回 Hello, World，得到 %q", got)
	}
}
