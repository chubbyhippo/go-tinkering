package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hippo")
	want := "Hello, Hippo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
