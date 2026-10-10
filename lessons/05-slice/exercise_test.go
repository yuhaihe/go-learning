package slicex

import "reflect"
import "testing"

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 2, 1, 3})
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Unique=%v，期望 %v", got, want)
	}
	if got := Unique(nil); len(got) != 0 {
		t.Fatalf("空输入应得到空 slice，得到 %v", got)
	}
}
