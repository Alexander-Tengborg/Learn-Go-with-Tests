package main

import (
	"os"
	"time"

	svg "github.com/alexander-tengborg/learn-go-with-tests/maths/clockface/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
