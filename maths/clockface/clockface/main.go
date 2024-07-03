package main

import (
	"os"
	"time"

	clockface "github.com/alexander-tengborg/learn-go-with-tests/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
