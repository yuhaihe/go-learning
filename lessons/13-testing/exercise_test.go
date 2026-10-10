package fake

import (
	"errors"
	"testing"
)

type fakeStore struct {
	names map[int]string
}

func (f fakeStore) Name(id int) (string, error) {
	name, ok := f.names[id]
	if !ok {
		return "", errors.New("not found")
	}
	return name, nil
}

func TestGreet(t *testing.T) {
	store := fakeStore{names: map[int]string{1: "ann"}}
	cases := []struct {
		id      int
		want    string
		wantErr bool
	}{
		{id: 1, want: "hi ann"},
		{id: 2, wantErr: true},
	}
	for _, tc := range cases {
		got, err := Greet(store, tc.id)
		if tc.wantErr {
			if err == nil {
				t.Fatalf("id %d 应返回错误", tc.id)
			}
			continue
		}
		if err != nil || got != tc.want {
			t.Fatalf("Greet(%d)=%q %v，期望 %q", tc.id, got, err, tc.want)
		}
	}
}
