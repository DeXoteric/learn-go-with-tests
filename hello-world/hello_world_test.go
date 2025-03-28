package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Marek")
	want := "Hello, Marek"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
