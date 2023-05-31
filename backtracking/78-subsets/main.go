package main

func main() {
	subsets([]int{1, 2, 3})
}

func subsets(nums []int) (res [][]int) {
	curr := make([]int, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		res = append(res, append([]int{}, curr...))
		if idx == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)
	return res
}
