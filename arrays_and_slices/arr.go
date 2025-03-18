package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbresToSum ...[]int) []int {
	lenghtOfNumbers := len(numbresToSum)
	sums := make([]int, lenghtOfNumbers)

	for i, numbers := range numbresToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
