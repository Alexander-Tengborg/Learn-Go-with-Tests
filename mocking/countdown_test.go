package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 2 1 Go! and sleeps 3 times at some point", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleepPrinter := SpyCountdownOperations{}

		Countdown(&buffer, &spySleepPrinter)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := SpyCountdownOperations{}

		Countdown(&spySleepPrinter, &spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spySleepPrinter.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter)
		}
	})
}
