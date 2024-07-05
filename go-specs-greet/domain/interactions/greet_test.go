package interactions

import (
	"testing"

	"github.com/alexander-tengborg/learn-go-with-tests/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(Greet),
	)
}
