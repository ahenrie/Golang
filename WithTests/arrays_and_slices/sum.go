package main

func Sum(numbers [5]int) int {
	total := 0
	for i := 0; i < 5; i++ {
		total += numbers[i]
	}
	return total
}
