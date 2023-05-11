package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0, len(nums)/3)

	if len(nums) < 3 {
		return res
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if target > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
