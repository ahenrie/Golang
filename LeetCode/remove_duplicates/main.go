package main

import "fmt"

func main() {
	nums := []int{1, 2, 2}
	fmt.Print(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	//map to keep track of ints that have been seen
	seen := make(map[int]bool)
	count := 0

	// range over nums
	for _, num := range nums {
		// if not in map
		if !seen[num] {
			//set to true in map
			seen[num] = true
			// modifies the nums array directly
			nums[count] = num
			count++
		}
	}
	return count
}
