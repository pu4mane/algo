package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	x := [3]int{3, 2, 1}

	Reverse(&x)

	got := x
	want := [3]int{1, 2, 3}

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
