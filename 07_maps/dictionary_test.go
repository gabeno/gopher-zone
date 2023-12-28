package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
