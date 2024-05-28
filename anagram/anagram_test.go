package anagram

import "testing"

func TestAnagram(t *testing.T) {
	t.Run("test true", func(t *testing.T) {
		src := "anagram"
		tgt := "nagaram"

		got := IsAnagram(src, tgt)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("test false", func(t *testing.T) {
		src := "anagram"
		tgt := "nagaraz"

		got := IsAnagram(src, tgt)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
