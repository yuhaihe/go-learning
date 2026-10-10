package shape

import "testing"

func TestArea(t *testing.T) {
	if got := (Rect{W: 3, H: 4}).Area(); got != 12 {
		t.Fatalf("面积=%d，期望 12", got)
	}
}

func TestLabel(t *testing.T) {
	u := User{Name: "ann", Address: Address{City: "hz"}}
	if got := u.Label(); got != "ann@hz" {
		t.Fatalf("Label=%q", got)
	}
}
