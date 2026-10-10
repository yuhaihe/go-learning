package mapsx

import "testing"

func TestCount(t *testing.T) {
	got := Count([]string{"go", "go", "test"})
	if got["go"] != 2 || got["test"] != 1 {
		t.Fatalf("计数错误：%v", got)
	}
	if _, ok := got["nope"]; ok {
		t.Fatal("不应出现未统计的键")
	}
}
