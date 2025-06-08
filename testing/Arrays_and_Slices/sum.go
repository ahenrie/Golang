package arr_slices

func sum(numbers []int) int {
	total := 0
	//for i := 0; i < len(numbers); i++ {
	for _, number := range numbers {
		total += number
	}
	return total
}
