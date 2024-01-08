package arrays


func sum(numbers []int) int {
	sum := 0

	for _, a := range numbers {
		sum += a
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int 

	for _ , numbers := range(numbersToSum) {
		sums = append(sums, sum(numbers))
	}
	return sums
}


func SumALLTails(numbersToSumTails ...[]int) []int {
	var tails []int
	for _, numbers := range(numbersToSumTails) {
		if len(numbers) == 0 {
			tails = append(tails, 0)
		} else {
		tail := numbers[1:]
		tails = append(tails, sum(tail))
		}

	}

	return tails
}