package main

func maxProfit(prices []int) int {
	min := prices[0]
	bestProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > bestProfit {
			bestProfit = prices[i] - min
		}
	}
	return bestProfit

}
