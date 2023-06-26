package main

func main() {
	maxProfit([]int{6, 1, 6, 4, 3, 0, 2})
}

func maxProfit(prices []int) int {
	sold, hold, rest := 0, -1001, 0

	for i := range prices {
		prevSold := sold

		sold = hold + prices[i]
		hold = max(hold, rest-prices[i])
		rest = max(rest, prevSold)
	}

	return max(sold, rest)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
