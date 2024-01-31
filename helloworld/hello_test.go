package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("empty string defaults to 'World", func(t *testing.T) {
		got := hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := hello("Antony", "French")
		want := "Bonjour, Antony"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
