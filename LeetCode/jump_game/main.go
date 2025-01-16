package main

func canJump(nums []int) bool {
	canMove := 0

	for i := 0; i < len(nums); i++ {
		if canMove < 0 {
			return false
		} else if nums[i] > canMove {
			canMove = nums[i]
		}
		canMove -= 1
	}
	return true
}
