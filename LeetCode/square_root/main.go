package main

import "math"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	n := float64(x)
	for i := 0; i < 100; i++ {
		n = (n + float64(x)/n) / 2
	}
	return int(math.Floor(n))
}
