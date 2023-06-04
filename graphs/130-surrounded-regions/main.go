package main

func main() {
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x == m || y == n || board[x][y] != 'O' {
			return
		}

		board[x][y] = '*'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}
