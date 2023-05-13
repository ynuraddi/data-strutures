package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures))

	for i, t := range temperatures {
		for j := len(stack) - 1; j >= 0 && temperatures[stack[j]] < t; j-- {
			res[stack[j]] = i - stack[j]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}
