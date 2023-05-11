package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	size := len(height)
	left, right := 0, size-1
	maxWidth := size - 1
	area := 0

	for width := maxWidth; width > 0; width-- {
		if height[left] < height[right] {
			area = max(area, width*height[left])
			left++
		} else {
			area = max(area, width*height[right])
			right--
		}
	}

	return area
}

func max(a, b int) int {
	if a <= b {
		return b
	}

	return a
}
