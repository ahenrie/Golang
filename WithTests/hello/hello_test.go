package main

import "testing"

func TestHello(t *testing.T) {
	// Subtest for a passing of name
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	// Subtest for no name
	t.Run("say 'Hello, World' when an empy string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// Subtest for spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("SomeName", "Spanish")
		want := "Hola, SomeName"
		assertCorrectMessage(t, got, want)
	})
}

// Helper function for comparison
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
