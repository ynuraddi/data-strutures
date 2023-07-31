package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
	fmt.Println(maxCoins([]int{1, 5}))
}

func maxCoins(nums []int) (result int) {
	nums = append([]int{1}, append(nums, 1)...)
	cache := make(map[[2]int]int)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if val, ok := cache[[2]int{i, j}]; ok {
			return val
		}

		out := 0
		for k := i + 1; k <= j-1; k++ {
			val := nums[i]*nums[k]*nums[j] + dfs(i, k) + dfs(k, j)
			if val > out {
				out = val
			}
		}
		cache[[2]int{i, j}] = out
		return out
	}

	return dfs(0, len(nums)-1)
}
