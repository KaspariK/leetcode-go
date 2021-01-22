package main

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	low := prices[0]
	maxProfit := 0

	for _, price := range prices {
		maxProfit = maxInt(maxProfit, price - low)

		if price < low {
			low = price
		}
	}

	return maxProfit
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
