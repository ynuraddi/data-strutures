package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numSubseq([]int{3, 5, 6, 7}, 9))
}

const mod int = 1e9 + 7

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	ans := 0

	pow2 := make([]int, len(nums)+1)
	pow2[0] = 1
	for i := 1; i <= len(nums); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	for left <= right {
		if nums[left]+nums[right] <= target {
			ans = (ans + pow2[right-left]) % mod
			left++
		} else {
			right--
		}
	}

	return ans
}
