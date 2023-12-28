package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("existing word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("non existing word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")

		assertError(t, err, ErrWordNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
