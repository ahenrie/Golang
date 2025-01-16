package main

import "fmt"

func majorityElement(nums []int) int {
	elementCounts := make(map[int]int)

	for _, num := range nums {
		elementCounts[num]++
	}

	var key int
	maxCount := 0

	//iterate over the map
	for num, count := range elementCounts {
		if count > maxCount {
			maxCount = count
			key = num
		}
	}
	return key
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7, 8, 5, 6, 7, 8, 8, 8, 8, 8, 8}
	fmt.Println(majorityElement(nums))
}
