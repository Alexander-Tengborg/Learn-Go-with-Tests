package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.141592653589793 * c.Radius 
}

func (c Circle) Area() float64 {
	return 3.141592653589793 * c.Radius * c.Radius
}