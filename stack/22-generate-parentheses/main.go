package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	stack := ""
	res := []string{}

	backtrack(stack, n, n, &res)

	return res
}

func backtrack(stack string, open int, closed int, res *[]string) {
	if open == 0 && closed == 0 {
		*res = append(*res, stack)
		return
	}
	if len(stack) == 0 || open != 0 {
		backtrack(stack+"(", open-1, closed, res)
	}
	if len(stack) != 0 && closed != 0 && closed > open {
		backtrack(stack+")", open, closed-1, res)
	}
}
