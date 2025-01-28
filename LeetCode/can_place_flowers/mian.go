package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	plantableSpots := 0

	for i := 0; i < len(flowerbed); {
		//Check if current spot is available
		if flowerbed[i] == 0 &&
			// Check left boundary
			(i == 0 || flowerbed[i-1] == 0) &&
			// check right boundary
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			plantableSpots++
			i += 2
		} else {
			i++
		}
	}

	return n <= plantableSpots
}
