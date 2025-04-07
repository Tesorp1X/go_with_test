package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("non-concurretn test", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 4)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d want %d", got, want)
	}
}
