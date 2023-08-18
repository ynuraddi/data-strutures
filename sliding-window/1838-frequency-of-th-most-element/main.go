package main

func main() {
	maxFrequency([]int{1, 4, 8, 13}, 5)
}

func maxFrequency(nums []int, k int) int {
	if len(nums) < 2 {
		return 1
	}

	var needed, ans int
	p1, p2 := 0, 0
	for p1 < len(nums) && p2 < len(nums) {
		if needed <= k {
			ans = max(ans, p2-p1+1)
			p2++
			if p2 < len(nums) {
				needed += (p2 - p1) * (nums[p2] - nums[p2-1])
			}
		} else {
			needed -= nums[p2] - nums[p1]
			p1++
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
