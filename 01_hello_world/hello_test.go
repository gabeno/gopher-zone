package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say hello with name", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello without name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Jao", "spanish")
		want := "Hola, Jao"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // tells test suite this method is a helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
