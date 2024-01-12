package romans

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRomans(1)
	want := "I"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
