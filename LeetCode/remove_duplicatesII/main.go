package main

func removeDuplicates(nums []int) int {
	twoTimes := make(map[int]int)
	count := 0

	for _, num := range nums {
		if twoTimes[num] < 2 {
			twoTimes[num]++
			nums[count] = num
			count++
		}
	}
	return count
}
