package intersection

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}

	got := Intersection(a, b)
	want := []int{2, 23}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
