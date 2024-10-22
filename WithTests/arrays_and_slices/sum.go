package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Variadic function to take 0 or many slices of ints
func SumAll(numbersToSum ...[]int) []int {
	// make a slice with length of number of input slices
	sums := make([]int, 0)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Variadic function to sum tails of slices
func SumAllTails(numbersToSum ...[]int) []int {
	tailSums := make([]int, 0)

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tail := numbers[1:]
			tailSums = append(tailSums, Sum(tail))
		}
	}

	return tailSums
}
