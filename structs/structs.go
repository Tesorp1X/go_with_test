package structs

const pi float64 = 3.1415

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimetr() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimetr() float64 {
	return 2 * pi * c.Radius
}

func (c Circle) Area() float64 {
	return pi * c.Radius * c.Radius
}
