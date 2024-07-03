package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func minutesInRadians(t time.Time) float64 {
	return (math.Pi / (minutesInHalfClock / float64(t.Minute()))) + secondsInRadians(t)/minutesInClock
}

func hoursInRadians(t time.Time) float64 {
	return (math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock))) + minutesInRadians(t)/hoursInClock
}

func simpleTime(hours int, minutes int, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
