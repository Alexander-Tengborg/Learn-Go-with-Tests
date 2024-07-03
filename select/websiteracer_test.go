package websiteracer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebookc.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}