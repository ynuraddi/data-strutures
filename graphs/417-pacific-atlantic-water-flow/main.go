package main

import "fmt"

func main() {
	fmt.Println(pacificAtlantic([][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}))
}

type flow struct {
	pacific  bool
	atlantic bool
}

func pacificAtlantic(heights [][]int) (res [][]int) {
	m, n := len(heights), len(heights[0])
	pac, atl := make(map[int]bool), make(map[int]bool)

	var dfs func(x, y int, visited map[int]bool, prev int)
	dfs = func(x, y int, visited map[int]bool, prev int) {
		if visited[x*n+y] ||
			x < 0 ||
			y < 0 ||
			x == m ||
			y == n ||
			heights[x][y] < prev {
			return
		}

		visited[x*n+y] = true
		dfs(x-1, y, visited, heights[x][y])
		dfs(x+1, y, visited, heights[x][y])
		dfs(x, y-1, visited, heights[x][y])
		dfs(x, y+1, visited, heights[x][y])
	}

	for x := 0; x < m; x++ {
		dfs(x, 0, pac, heights[x][0])
		dfs(x, n-1, atl, heights[x][n-1])
	}

	for y := 0; y < n; y++ {
		dfs(0, y, pac, heights[0][y])
		dfs(m-1, y, atl, heights[m-1][y])
	}

	for i := range heights {
		for j := range heights[i] {
			if pac[i*n+j] && atl[i*n+j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return
}
