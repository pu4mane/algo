package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 8, 2, 12, 1, 10, 76, 10}

	got := BubbleSort(arr)
	want := []int{1, 2, 5, 8, 10, 10, 12, 76}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
