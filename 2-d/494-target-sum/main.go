package main

func main() {
	findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
}

func findTargetSumWays(nums []int, target int) int {
	caching := make(map[[2]int]int)

	var dfs func(idx, sum int) int
	dfs = func(idx, sum int) int {
		if val, ok := caching[[2]int{idx, sum}]; ok {
			return val
		}

		if idx == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		res := dfs(idx+1, sum+nums[idx]) + dfs(idx+1, sum-nums[idx])

		caching[[2]int{idx, sum}] = res
		return res
	}
	return dfs(0, 0)
}
