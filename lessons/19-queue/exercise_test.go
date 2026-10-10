package taskx

import (
	"encoding/json"
	"testing"
)

func TestNewNoteTask(t *testing.T) {
	task, err := NewNoteTask(3, "hello")
	if err != nil {
		t.Fatal(err)
	}
	if task.Type() != TaskNoteCreated {
		t.Fatalf("类型 %s", task.Type())
	}
	var body struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
	}
	if err := json.Unmarshal(task.Payload(), &body); err != nil {
		t.Fatal(err)
	}
	if body.ID != 3 || body.Title != "hello" {
		t.Fatalf("载荷 %+v", body)
	}
}
