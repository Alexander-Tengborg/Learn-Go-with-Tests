package main

import (
	"log"
	"net/http"

	httpserver "github.com/alexander-tengborg/learn-go-with-tests/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
