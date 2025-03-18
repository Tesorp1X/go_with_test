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
*/
func assertCorrect(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaTT(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{10.0, 20.0}, want: 200.0},
		{shape: Circle{10.0}, want: 10 * 10 * math.Pi},
		{shape: Triangle{4.0, 3.0, 5.0}, want: 6.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertCorrect(t, got, tt.want)
	}
}

func TestPerimetrTT(t *testing.T) {
	perimetrTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{10.0, 20.0}, want: 60.0},
		{shape: Circle{10.0}, want: 2 * 10 * math.Pi},
		{shape: Triangle{4.0, 3.0, 5.0}, want: 12.0},
	}

	for _, tt := range perimetrTests {
		got := tt.shape.Perimetr()
		assertCorrect(t, got, tt.want)
	}
}
