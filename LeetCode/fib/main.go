package main

import "fmt"

func fib(n int) int {
	nums := []int{0, 1}

	for i := 2; i < n+2; i++ {
		nums = append(nums, nums[i-2]+nums[i-1])
	}
	return nums[n]
}

func main() {
	fmt.Println(fib(20))
}
