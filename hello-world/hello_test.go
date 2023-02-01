package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Hippo", "")
		want := "Hello, Hippo"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, Hippo' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Hippo"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Hippo", "Spanish")
		want := "Hola, Hippo"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
