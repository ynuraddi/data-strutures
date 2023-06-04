package main

func main() {
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	numIslands := 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m ||
			y < 0 || y >= n ||
			grid[x][y] == '0' {
			return
		}

		grid[x][y] = '0'

		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	for i, g := range grid {
		for j, point := range g {
			if point == '1' {
				dfs(i, j)
				numIslands += 1
			}
		}
	}
	return numIslands
}
