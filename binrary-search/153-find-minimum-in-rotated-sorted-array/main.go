package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3, 4}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	res := nums[0]
	for l <= r {
		if nums[l] < nums[r] {
			res = min(res, nums[l])
		}

		mid := (l + r) / 2
		res = min(res, nums[mid])
		if nums[mid] >= nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
