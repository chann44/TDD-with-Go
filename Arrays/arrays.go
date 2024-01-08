package arrays


func sum(numbers []int) int {
	sum := 0

	for _, a := range numbers {
		sum += a
	}

	return sum
}