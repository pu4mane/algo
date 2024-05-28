package countvowels

import "testing"

func TestCountVowels(t *testing.T) {
	t.Run("test true", func(t *testing.T) {
		str := "andrey"

		got := CountVowels(str)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
