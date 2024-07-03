package websiteracer

import (
	"net/http"
	"time"
)

func Racer(a string, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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