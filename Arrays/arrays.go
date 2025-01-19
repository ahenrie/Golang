package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	//create an array that will hold 5 ints
	var a [5]int
	fmt.Println("emp:", a)

	//all indices will be written with a 0

	//access 5th int
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//print length of an array
	fmt.Println("len:", len(a))

	//declare and initialize in one line
	b := [5]int{1, 2, 3, 4}

	//same but use the compiler to count
	//b = [...]int{0, 0:100}
	fmt.Println(b)

	// Construct a sorted count array
	// I know it is only going to be of length 10 since my numbers range from 0 to 10
	slice1 := []int{}
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		slice1 = append(slice1, rand.Intn(10))
	}

	countArr := make([]int, 11)
	for _, num := range slice1 {
		countArr[num]++
	}
	fmt.Println(countArr)

	// Make the countArr a zigzag
	n := len(countArr)
	sort.Ints(countArr)
	m := (n - 1) / 2
	fmt.Println(countArr)
	for i, j := m, n-1; i < j; i, j = i+1, j-1 {
		countArr[i], countArr[j] = countArr[j], countArr[i]
	}
	fmt.Println(countArr)

	//Remove duplicates
	set := make(map[int]bool)

	for _, num := range countArr {
		set[num] = true
	}

	uniqueNums := []int{}
	for key, _ := range set {
		uniqueNums = append(uniqueNums, key)
	}

	sort.Ints(uniqueNums)
	fmt.Println(uniqueNums)

}
