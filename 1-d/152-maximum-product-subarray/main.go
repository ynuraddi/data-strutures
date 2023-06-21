package main

func main() {
	maxProduct([]int{2, 3, -2, 4})
}

func maxProduct(nums []int) int {
	res, currMin, currMax := nums[0], 1, 1

	for i := 0; i < len(nums); i++ {
		tmp := currMax * nums[i]
		currMax = max(max(nums[i]*currMax, nums[i]*currMin), nums[i])
		currMin = min(min(tmp, currMin*nums[i]), nums[i])
		res = max(res, currMax)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
