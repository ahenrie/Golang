package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	greatestCandies := []bool{}

	for i := 0; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxCandies {
			greatestCandies = append(greatestCandies, true)
		} else {
			greatestCandies = append(greatestCandies, false)
		}
	}

	return greatestCandies

}
