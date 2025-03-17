package arraysandslices

func Sum(numbers [3]int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}
