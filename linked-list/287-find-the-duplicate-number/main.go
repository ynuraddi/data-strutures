package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 7, 5, 2}))
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}
