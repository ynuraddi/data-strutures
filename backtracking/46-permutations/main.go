package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) (res [][]int) {
	lenNums := len(nums)
	var bfs func(curr, used []int)
	bfs = func(curr, used []int) {
		if len(used) == lenNums {
			res = append(res, used)
			return
		}

		for i := 0; i < len(curr); i++ {
			bfs(append(append([]int{}, curr[:i]...), curr[i+1:]...), append(used, curr[i]))
		}
	}
	bfs(nums, []int{})
	return res
}
