package websiteracer

import (
	"net/http"
	"time"
)

func Racer(a string, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)
	
	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	duration := time.Since(startTime)

	return duration
}