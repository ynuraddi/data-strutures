package main

func main() {
	coinChange([]int{186, 419, 83, 408}, 6249)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := range dp {
		for _, c := range coins {
			if c <= i {
				dp[i] = min(dp[i], 1+dp[i-c])
			}
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
