package arraysandslices

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("Sum 2 slices separatly", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 9})
		expected := []int{3, 12}

		assertCorrect(t, got, expected)
	})
	t.Run("Sum 3 slises separatly", func(t *testing.T) {
		got := SumAll([]int{10, 2}, []int{3, 9, -1}, []int{1, 2, 3, 4})
		expected := []int{12, 11, 10}

		assertCorrect(t, got, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("slices of length 2", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		assertCorrect(t, got, expected)
	})
	t.Run("slices of any length", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{0, 9, 1})
		expected := []int{9, 10}
		assertCorrect(t, got, expected)
	})
	t.Run("slices of length 1", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{0, 9})
		expected := []int{0, 9}
		assertCorrect(t, got, expected)
	})
}

func assertCorrect(t testing.TB, got, expected []int) {
	if !slices.Equal(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}
