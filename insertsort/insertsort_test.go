package insertsort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{5, 8, 2, 12, 1, 10, 76, 10}

	got := InsertSort(arr)
	want := []int{1, 2, 5, 8, 10, 10, 12, 76}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
