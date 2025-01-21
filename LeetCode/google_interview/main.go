package main

import (
	"fmt"
	"math/rand"
)

/*
Design a class that:
1. Inserting a value (no duplicates allowed)
2. Removing a value
3. GetRandom --> returns a random value with equal probability of returning a value that is already inserted.
*/

type CoolClass struct {
	//Data struct that stores unique values, remove, and getRandom (hopefully at O(1))
	intMap map[int]bool
	nums   []int
}

func newCoolClass() *CoolClass {
	return &CoolClass{
		intMap: make(map[int]bool),
		nums:   []int{},
	}
}

// Insert (NO DUPLICATES) (O1)
func (cc *CoolClass) insert(v int) bool {
	if !cc.intMap[v] {
		cc.intMap[v] = true
		cc.nums = append(cc.nums, v)
		return true
	} else {
		return false
	}
}

func (cc *CoolClass) remove(v int) bool {
	if !cc.intMap[v] {
		return false
	} else {
		// set to false in the intMap
		cc.intMap[v] = false
		// find the number then swap it out with the last, then pop nums
		for i, num := range cc.nums {
			if num == v {
				// assign last num to where num to remove is then pop off the last num to remove duplicates
				cc.nums[i] = cc.nums[len(cc.nums)-1]
				cc.nums = cc.nums[:len(cc.nums)-1]
				return true
			}
		}
	}
	return false
}

// Return a random int with equal probability
func (cc *CoolClass) getRandom() int {
	// Check length
	if len(cc.nums) == 0 {
		return -1
	}

	//rand.Seed(time.Now().UnixNano())
	return cc.nums[rand.Intn(len(cc.nums))]
}
func main() {
	testClass := newCoolClass()
	testClass.insert(1)
	testClass.insert(2)
	testClass.insert(3)
	testClass.insert(4)
	testClass.insert(5)

	testClass.remove(2)

	randomInts := []int{}
	for i := 0; i < 100; i++ {
		randomInts = append(randomInts, testClass.getRandom())
	}
	fmt.Println(randomInts)
}
