package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Linda")
	want := "Hello, Linda"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
