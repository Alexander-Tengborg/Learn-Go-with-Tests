package websiteracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares the speed of 2 servers and returns the URL of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Millisecond * 20)
		fastServer := makeDelayedServer(time.Millisecond * 0)
		
		defer slowServer.Close()
		defer fastServer.Close()
		
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		
		want := fastURL
		got, err := Racer(slowURL, fastURL)
		
		if err != nil {
			t.Fatalf("should not have gotten an error")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the given timeout time", func(t *testing.T) {
		delayedServer := makeDelayedServer(time.Millisecond * 50)
		
		defer delayedServer.Close()
		
		delayedURL := delayedServer.URL
		
		_, err := ConfigurableRacer(delayedURL, delayedURL, time.Millisecond * 20)
		
		if err == nil {
			t.Errorf("should have gotten an error from a server not responding within 10s")
		}	
	})
}

func makeDelayedServer(sleepTime time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepTime)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}