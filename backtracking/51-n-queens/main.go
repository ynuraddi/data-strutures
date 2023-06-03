package main

func main() {
	solveNQueens(4)
}

func solveNQueens(n int) (res [][]string) {
	curr := make([]string, 0)
	column, d1, d2 := make(map[int]int), make(map[int]int), make(map[int]int)
	var bfs func(y int)
	bfs = func(y int) {
		if y == n {
			res = append(res, append([]string{}, curr...))
		}

		for x := 0; x < n; x++ {
			if column[x] > 0 || d1[x+y] > 0 || d2[n+x-y-1] > 0 {
				continue
			}
			column[x], d1[x+y], d2[n+x-y-1] = 1, 1, 1
			s := ""
			for i := 0; i < n; i++ {
				if i == x {
					s += "Q"
				} else {
					s += "."
				}
			}
			curr = append(curr, s)
			bfs(y + 1)
			column[x], d1[x+y], d2[n+x-y-1] = 0, 0, 0
			curr = curr[:len(curr)-1]
		}
	}
	bfs(0)
	return
}
