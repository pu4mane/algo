package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	n := 7

	got := Fibonacci(n)
	want := 13

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
