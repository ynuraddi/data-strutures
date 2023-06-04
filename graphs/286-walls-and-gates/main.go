package main

func main() {
	WallsAndGates([][]int{
		{-2, -1, 0, -2},
		{-2, -2, -2, -1},
		{-2, -1, -2, -1},
		{0, -1, -2, -2},
	})
}

func WallsAndGates(rooms [][]int) {
	m, n := len(rooms), len(rooms[0])

	positions := make([][]int, 0)
	roomCount := 0

	for i := range rooms {
		for j := range rooms[i] {
			if rooms[i][j] == 0 {
				positions = append(positions, []int{i, j})
			}
			if rooms[i][j] < -1 {
				roomCount += 1
			}
		}
	}

	dfs := func(x, y, distance int) {
		if x < 0 || y < 0 || x == m || y == n {
			return
		}
		if rooms[x][y] < -1 {
			positions = append(positions, []int{x, y})
			rooms[x][y] = distance
			roomCount -= 1
		}
	}

	distance := 0
	tick := len(positions)
	for roomCount != 0 && len(positions) > 0 {
		distance += 1
		for i := 0; i < tick; i++ {
			x, y := positions[0][0], positions[0][1]

			dfs(x+1, y, distance)
			dfs(x-1, y, distance)
			dfs(x, y+1, distance)
			dfs(x, y-1, distance)

			positions = positions[1:]
		}
		tick = len(positions)
	}
}
