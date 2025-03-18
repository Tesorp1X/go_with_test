package structs

import (
	"math"
	"testing"
)

/*
Depricated tests

	func TestPerimetr(t *testing.T) {
		t.Run("rectangle", func(t *testing.T) {
			fig := Rectangle{10.0, 20.0}
			got := fig.Perimetr()
			want := 60.0

			assertCorrect(t, got, want)
		})
		t.Run("circle", func(t *testing.T) {
			fig := Circle{10.0}
			got := fig.Perimetr()
			want := 2 * 10 * math.Pi

			assertCorrect(t, got, want)
		})
	}

	func TestArea(t *testing.T) {
		t.Run("rectangle", func(t *testing.T) {
			fig := Rectangle{10.0, 20.0}
			got := fig.Area()
			want := 200.0

			assertCorrect(t, got, want)
		})
		t.Run("circle", func(t *testing.T) {
			fig := Circle{10.0}
			got := fig.Area()
			want := 10 * 10 * math.Pi

			assertCorrect(t, got, want)
		})
	}

	func assertCorrect(t testing.TB, got, want float64) {
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
*/
func TestAreaTT(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 20.0}, hasArea: 200.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 10 * 10 * math.Pi},
		{name: "Triangle", shape: Triangle{a: 4.0, b: 3.0, c: 5.0}, hasArea: 6.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestPerimetrTT(t *testing.T) {
	perimetrTests := []struct {
		name        string
		shape       Shape
		hasPerimetr float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 20.0}, hasPerimetr: 60.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasPerimetr: 2 * 10 * math.Pi},
		{name: "Triangle", shape: Triangle{a: 4.0, b: 3.0, c: 5.0}, hasPerimetr: 12.0},
	}

	for _, tt := range perimetrTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimetr()
			if got != tt.hasPerimetr {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasPerimetr)
			}
		})

	}
}
