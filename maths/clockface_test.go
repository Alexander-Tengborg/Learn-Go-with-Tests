package clockface

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), math.Pi / 2 * 3},
		{simpleTime(0, 0, 7), math.Pi / 30 * 7},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("converting %vs to %v radians", test.time.Second(), test.angle), func(t *testing.T) {
			got := secondsInRadians(test.time)

			if test.angle != got {
				t.Fatalf("wanted %v radians, got %v radians", test.angle, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("converting %vs to point %v", test.time.Second(), test.point), func(t *testing.T) {
			got := secondHandPoint(test.time)

			if !roughlyEqualPoint(test.point, got) {
				t.Fatalf("wanted %v, got %v", test.point, got)
			}
		})
	}
}

func roughlyEqualFloat64(a float64, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a Point, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
