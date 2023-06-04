package main

func main() {
	orangesRotting([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}) // -1
	orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}) // 4
}

func orangesRotting(grid [][]int) (timer int) {
	n := len(grid[0])

	fresh := make(map[int]struct{})
	rotten := make([]int, 0)

	freshCount := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fresh[i*n+j] = struct{}{}
				freshCount += 1
			}
			if grid[i][j] == 2 {
				rotten = append(rotten, i*n+j)
			}
		}
	}

	dfs := func(position int) {
		if _, ok := fresh[position]; ok {
			rotten = append(rotten, position)
			delete(fresh, position)
			freshCount -= 1
		}
	}

	idx := len(rotten)
	for len(rotten) != 0 && freshCount != 0 {
		timer += 1
		for i := 0; i < idx; i++ {
			rotPosition := rotten[0]

			dfs(rotPosition - n)
			dfs(rotPosition + n)
			if (rotPosition+n)%n != 0 {
				dfs(rotPosition - 1)
			}
			if (rotPosition+1)%n != 0 {
				dfs(rotPosition + 1)
			}

			rotten = rotten[1:]
		}
		idx = len(rotten)
	}

	if freshCount != 0 {
		return -1
	}
	return timer
}
