package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 3 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})

	t.Run("Collection of any length", func(t *testing.T) {
		numberSlice := []int{1, 2, 3, 4}

		got := Sum(numberSlice)
		expected := 10

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numberSlice)
		}
	})
}
