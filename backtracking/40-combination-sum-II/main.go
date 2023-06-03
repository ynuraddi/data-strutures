package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	curr := make([]int, 0)
	sum := 0
	var bfs func(idx int)
	bfs = func(idx int) {
		if sum == target {
			res = append(res, append([]int{}, curr...))
			return
		}
		if sum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			curr = append(curr, candidates[i])
			bfs(i + 1)
			curr = curr[:len(curr)-1]
			sum -= candidates[i]
		}
	}
	bfs(0)
	return
}
