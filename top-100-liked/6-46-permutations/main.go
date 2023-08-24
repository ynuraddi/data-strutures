package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	result := make([][]int, 0, len(nums))

	var bfs func(curr, mem []int)
	bfs = func(curr, mem []int) {
		if len(mem) == len(nums) {
			result = append(result, mem)
			return
		}

		for i, v := range curr {
			bfs(
				append(append([]int{}, curr[:i]...), curr[i+1:]...),
				append(mem, v),
			)
		}
	}

	bfs(nums, []int{})
	return result
}
