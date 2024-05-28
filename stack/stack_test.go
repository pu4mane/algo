package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("test Push", func(t *testing.T) {
		s := Stack{[]int{2, 4}}

		item := 4

		s.Push(item)

		got := s.items[len(s.items)-1]
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("test Pop", func(t *testing.T) {
		s := Stack{[]int{2, 4, 3}}

		s.Pop()

		got := len(s.items)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("test Top", func(t *testing.T) {
		s := Stack{[]int{2, 4, 3}}

		got, _ := s.Top()
		want := 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
