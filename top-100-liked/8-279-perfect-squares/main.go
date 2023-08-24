package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(13))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt - 10
	}
	dp[0] = 0

	nums := make([]int, 0, 10)
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	for i := range dp {
		for _, v := range nums {
			if i-v >= 0 {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}

	fmt.Println(dp)
	fmt.Println(nums)
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
