package main

func main() {
	maxAreaOfIsland([][]int{})
}

func maxAreaOfIsland(grid [][]int) (maxArea int) {
	m, n := len(grid), len(grid[0])
	tmp := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m ||
			y < 0 || y >= n ||
			grid[x][y] == 0 {
			return
		}

		tmp += 1
		grid[x][y] = 0

		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for x, g := range grid {
		for y, point := range g {
			if point != 0 {
				dfs(x, y)
				if tmp > maxArea {
					maxArea = tmp
				}
				tmp = 0
			}
		}
	}
	return
}
