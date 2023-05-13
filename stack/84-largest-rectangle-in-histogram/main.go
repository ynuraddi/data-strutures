package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

type Bar struct {
	Index  int
	Height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	ret := 0
	size := len(heights)
	stack := []Bar{{-1, 0}}

	for i := 0; i < size; i++ {
		last := len(stack) - 1
		height := stack[last].Height

		if heights[i] < height {
			width := i - stack[last-1].Index - 1

			ret = max(ret, height*width)
			stack = stack[:last]
			i--
		} else if heights[i] > height {
			stack = append(stack, Bar{i, heights[i]})
		} else {
			stack[last].Index = i
		}
	}

	for i := len(stack) - 1; i > 0; i-- {
		height := stack[i].Height
		width := size - stack[i-1].Index - 1

		ret = max(ret, height*width)
	}

	return ret
}
