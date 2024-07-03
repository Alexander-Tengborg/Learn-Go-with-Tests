package romannumerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("converting 1 to roman numerals", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}