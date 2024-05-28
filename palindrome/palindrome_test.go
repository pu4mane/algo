package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	t.Run("test true", func(t *testing.T) {
		str := "anna"

		got := IsPalindrome(str)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("test false", func(t *testing.T) {
		str := "drey"

		got := IsPalindrome(str)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
