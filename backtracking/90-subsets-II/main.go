package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) (res [][]int) {
	sort.Ints(nums)
	n := len(nums)
	curr := make([]int, 0)

	var bfs func(idx int)
	bfs = func(idx int) {
		res = append(res, append([]int{}, curr...))
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			bfs(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	bfs(0)
	return
}
