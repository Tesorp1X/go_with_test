package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Array of three elements", func(t *testing.T) {
		numbers := [3]int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})

}
