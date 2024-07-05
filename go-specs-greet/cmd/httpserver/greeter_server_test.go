package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/alexander-tengborg/learn-go-with-tests/go-specs-greet/adapters"
	"github.com/alexander-tengborg/learn-go-with-tests/go-specs-greet/adapters/httpserver"
	"github.com/alexander-tengborg/learn-go-with-tests/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./go-specs-greet/cmd/httpserver/Dockerfile"
		BaseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{
			BaseURL: BaseURL,
			Client: &http.Client{
				Timeout: 1 * time.Second,
			},
		}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
