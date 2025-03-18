package structs

import "testing"

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
		want := 2 * 10 * pi

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
		want := 10 * 10 * pi

		assertCorrect(t, got, want)
	})
}

func assertCorrect(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
