package repeat

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat string characters 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})
}
