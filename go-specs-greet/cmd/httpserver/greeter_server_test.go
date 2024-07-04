package main_test

import (
	"testing"

	go_specs_greet "github.com/alexander-tengborg/learn-go-with-tests/go-specs-greet"
	"github.com/alexander-tengborg/learn-go-with-tests/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
