package main

func moveZeroes(nums []int) {
	write := 0 // Pointer to track the position to write non-zero values

	// Move all non-zero elements forward
	for _, num := range nums {
		if num != 0 {
			nums[write] = num
			write++
		}
	}

	// Fill the remaining elements with zeros
	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}
}
