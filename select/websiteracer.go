package websiteracer

import (
	"fmt"
	"net/http"
	"time"
	// "time"
)

var tenSecondTimeout = time.Second * 10

func Racer(a string, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a string, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})

	go func() {
		http.Get(url)
		close(channel)
	}()

	return channel
}