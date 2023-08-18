package main

func main() {
	jump([]int{2, 3, 1, 1, 4})
}

func jump(nums []int) int {
	dp := make([]int, len(nums))

	for i := range dp {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			if dp[j] == 0 {
				dp[j] = dp[i] + 1
				continue
			}
			dp[j] = min(dp[i]+1, dp[j])
		}
	}

	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
