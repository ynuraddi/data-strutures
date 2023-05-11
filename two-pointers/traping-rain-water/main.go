package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	maxleft := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if i > 0 && maxleft[i-1] <= height[i] {
			maxleft[i] = height[i]
		} else if i > 0 {
			maxleft[i] = maxleft[i-1]
		} else {
			maxleft[i] = height[i]
		}
	}

	maxright := make([]int, len(height))
	for i := len(height) - 1; i > 0; i-- {
		if i < len(height)-1 && maxright[i+1] <= height[i] {
			maxright[i] = height[i]
		} else if i < len(height)-1 {
			maxright[i] = maxright[i+1]
		} else {
			maxright[i] = height[i]
		}
	}

	res := 0
	for i := 0; i < len(height); i++ {
		sum := min(maxleft[i], maxright[i]) - height[i]
		if sum > 0 {
			res += sum
		}
	}

	return res
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
