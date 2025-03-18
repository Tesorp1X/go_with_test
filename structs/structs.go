package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimetr() float64 {
	return 2 * (r.Height + r.Width)
}
