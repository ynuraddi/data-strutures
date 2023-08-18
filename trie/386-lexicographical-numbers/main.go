package main

func main() {
}

func lexicalOrder(n int) []int {
	res := make([]int, 0, n)

	var dfs func(cur, limit int)
	dfs = func(cur, limit int) {
		if cur > limit {
			return
		}
		res = append(res, cur)
		for i := 0; i <= 9; i++ {
			if cur*10+i > limit {
				return
			}
			dfs(cur*10+i, limit)
		}
	}

	for i := 1; i <= 9; i++ {
		dfs(i, n)
	}
	return res
}
