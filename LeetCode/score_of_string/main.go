package main

import (
	"fmt"
	"math"
	//"strconv"
)

func scoreString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		a := int(s[i])
		b := int(s[i+1])
		sum += int(math.Abs(float64(a - b)))
	}
	return sum
}

func main() {
	s := "zaz"
	fmt.Println(scoreString(s))
}
