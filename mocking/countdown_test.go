package main

import (
	"testing"
	"bytes"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	SpySleeper := SpySleeper{}

	Countdown(&buffer, &SpySleeper)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if SpySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", SpySleeper.Calls)
	}
}