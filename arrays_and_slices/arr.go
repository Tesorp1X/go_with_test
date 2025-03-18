package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbresToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbresToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbresToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbresToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
