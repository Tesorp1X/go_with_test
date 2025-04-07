package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("non-concurrent Counter test", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 4)
	})

	t.Run("concurrent Counter test", func(t *testing.T) {
		wantedCount := 1000000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d want %d", got, want)
	}
}
