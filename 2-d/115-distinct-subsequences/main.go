package main

import "fmt"

func main() {
	a := numDistinct("rabbbit", "rabbit")
	fmt.Println(a)
}

func numDistinct(s string, t string) int {
	if t == "" {
		return 1
	} else if s == "" {
		return 0
	}

	cache := make(map[int]int, len(s))
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == len(t) {
			return 1
		} else if i == len(s) {
			return 0
		} else if v, ok := cache[i*len(t)+j]; ok {
			return v
		}

		var res int
		if s[i] == t[j] {
			res = dfs(i+1, j+1) + dfs(i+1, j)
		} else {
			res = dfs(i+1, j)
		}
		cache[i*len(t)+j] = res
		return res
	}
	return dfs(0, 0)
}
