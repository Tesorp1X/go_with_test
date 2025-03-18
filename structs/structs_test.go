package structs

import "testing"

func TestPerimetr(t *testing.T) {
	fig := Rectangle{10.0, 20.0}
	got := fig.Perimetr()
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
