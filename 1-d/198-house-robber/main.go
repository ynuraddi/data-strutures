package main

func main() {
	rob([]int{1, 2, 3, 1})
}

func rob(nums []int) (sum int) {
	dp := make([]int, len(nums)+2)

	for i := 2; i < len(dp); i++ {
		if dp[i-1] > dp[i-2]+nums[i-2] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i-2]
		}
	}

	return dp[len(dp)-1]
}
