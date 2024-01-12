package romans

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("1 converts to I", func(t *testing.T) {
		got := ConvertToRomans(1)
		want := "I"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("2 converts to II", func(t *testing.T) {
		got := ConvertToRomans(2)
		want := "II"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
