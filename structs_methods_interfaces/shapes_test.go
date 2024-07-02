package structsmethodsinterfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		got := shape.Perimeter()
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		want := 40.0
		
		checkPerimeter(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 2 * math.Pi * 10
		
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{shape: Rectangle{Width: 10, Height: 10}, want: 100.0},
		{shape: Circle{Radius: 10}, want: math.Pi * 100.0},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		
		if got != test.want {
			t.Errorf("got %g want %g", got, test.want)
		}
	}
}
