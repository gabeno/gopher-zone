package dictionary

import "testing"

func TestSearch(t *testing.T) {
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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")
	got, err := dictionary.Search("test")
	want := "this is a test"

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, want)
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
