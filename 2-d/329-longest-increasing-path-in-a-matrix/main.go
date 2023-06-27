package main

func main() {
	longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	})
}

func longestIncreasingPath(matrix [][]int) (res int) {
	m, n := len(matrix), len(matrix[0])
	cache := make(map[int]int, m*n)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if val, ok := cache[i*n+j]; ok {
			return val
		}

		var res int
		if i+1 < m && matrix[i][j] < matrix[i+1][j] {
			res = max(res, dfs(i+1, j))
		}
		if i-1 >= 0 && matrix[i][j] < matrix[i-1][j] {
			res = max(res, dfs(i-1, j))
		}
		if j+1 < n && matrix[i][j] < matrix[i][j+1] {
			res = max(res, dfs(i, j+1))
		}
		if j-1 >= 0 && matrix[i][j] < matrix[i][j-1] {
			res = max(res, dfs(i, j-1))
		}

		res += 1
		cache[i*n+j] = res
		return res
	}

	for i := range matrix {
		for j := range matrix[i] {
			if val := dfs(i, j); val > res {
				res = val
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
