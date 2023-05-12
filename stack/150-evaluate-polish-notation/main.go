package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}

func evalRPN(tokens []string) int {
	stack := stack{make([]int, 0)}
	operations := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}

	for _, n := range tokens {
		if _, ok := operations[n]; ok {
			n1, n2 := stack.poptop(), stack.poptop()
			switch n {
			case "+":
				stack.push(n2 + n1)
			case "-":
				stack.push(n2 - n1)
			case "*":
				stack.push(n2 * n1)
			case "/":
				stack.push(n2 / n1)
			}
		} else {
			num, _ := strconv.Atoi(n)
			stack.push(num)
		}
	}

	return stack.poptop()
}

type stack struct {
	nums []int
}

func (s *stack) push(val int) {
	s.nums = append(s.nums, val)
}

func (s *stack) poptop() int {
	num := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return num
}

// func (s *stack) len() int {
// 	return len(s.nums)
// }
