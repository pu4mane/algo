package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 4, 10, 12, 44, 876, 1000}
	target := 1

	got := BinarySearch(nums, target)
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
