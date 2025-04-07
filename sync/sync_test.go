package sync

import "testing"

func TestCounter(t *testing.T) {
	counter := Counter{}
	counter.Inc()
	counter.Inc()
	counter.Inc()
	counter.Inc()

	got := counter.Value()
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
