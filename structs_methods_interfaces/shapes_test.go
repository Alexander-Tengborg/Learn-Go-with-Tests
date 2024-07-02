package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := rectangle.Perimeter()
		want := 40.0
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 2 * 3.141592653589793 * 10
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := rectangle.Area()
		want := 100.0
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
