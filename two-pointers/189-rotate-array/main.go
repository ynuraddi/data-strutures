package main

import "fmt"

func main() {
	mass := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(mass, 2)
	fmt.Println(mass)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
