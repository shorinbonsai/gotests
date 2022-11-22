package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying heelo to peeps", func(t *testing.T) {
		got := Hello("James", "")
		want := "Hello, James"
		assertCorrect(t, got, want)
	})
	t.Run("empty string default to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrect(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("James", spanish)
		want := "Hola, James"
		assertCorrect(t, got, want)
	})
	t.Run("in de French", func(t *testing.T) {
		got := Hello("James", french)
		want := "Bonjour, James"
		assertCorrect(t, got, want)
	})

}

func assertCorrect(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
