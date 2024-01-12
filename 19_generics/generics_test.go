package generics_test

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("assert on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("assert on strings", func(t *testing.T) {
		AssertEqual(t, "gabe", "gabe")
		AssertNotEqual(t, "gabe", "gobe")
	})
}

func AssertEqual(t testing.TB, got, want interface{}) {
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func AssertNotEqual(t testing.TB, got, want interface{}) {
	if got == want {
		t.Errorf("did not want %+v", got)
	}
}
