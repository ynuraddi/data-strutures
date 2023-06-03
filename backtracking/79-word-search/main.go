package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB"))
	exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")
}

func exist(board [][]byte, word string) bool {
	ROWS, COLS := len(board), len(board[0])
	var bfs func(x, y, idx int) bool
	bfs = func(x, y, idx int) bool {
		if x < 0 || y < 0 ||
			x >= ROWS || y >= COLS ||
			idx == len(word) ||
			board[x][y] == '*' ||
			board[x][y] != word[idx] {
			return false
		}

		if idx == len(word)-1 {
			return true
		}

		tmp := board[x][y]
		board[x][y] = '*'

		res := bfs(x+1, y, idx+1) || bfs(x-1, y, idx+1) || bfs(x, y+1, idx+1) || bfs(x, y-1, idx+1)

		board[x][y] = tmp
		return res
	}
	for x, b := range board {
		for y := range b {
			if bfs(x, y, 0) {
				return true
			}
		}
	}
	return false
}
