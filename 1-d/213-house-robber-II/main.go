package main

func main() {
	rob([]int{1, 2, 3, 1})
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if a, b := robber(nums[1:]), robber(nums[:len(nums)-1]); a > b {
		return a
	} else {
		return b
	}
}

func robber(nums []int) int {
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
