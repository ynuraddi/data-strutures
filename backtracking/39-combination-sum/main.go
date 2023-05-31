package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) (res [][]int) {
	var bfs func([]int, int, int)
	bfs = func(nums []int, sum, idx int) {
		if sum == target {
			res = append(res, append([]int{}, nums...))
			return
		}
		if sum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			nums = append(nums, candidates[i])
			bfs(nums, sum+candidates[i], i)
			nums = nums[:len(nums)-1]
		}
	}
	bfs([]int{}, 0, 0)
	return res
}
