package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	slice1 := nums1[:m]
	slice2 := nums2[:n]
	slice1 = append(slice1, slice2...)
	sort.Ints(slice1)
	fmt.Println(slice1)
}

func main() {
	arr1 := [6]int{1, 2, 3, 0, 0, 0}
	arr2 := [3]int{2, 5, 6}
	m := 3
	n := 3
	merge(arr1[:], m, arr2[:], n)
}
