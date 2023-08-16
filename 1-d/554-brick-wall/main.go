package main

import "fmt"

func main() {
	fmt.Println(leastBricks([][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}))
}

func leastBricks(wall [][]int) int {
	hash := make(map[int]int, len(wall))

	var max int

	for i := range wall {
		currWhole := 0
		for j, brick := range wall[i] {
			if j == len(wall[i])-1 {
				continue
			}
			currWhole += brick
			hash[currWhole] += 1
			max = maxim(max, hash[currWhole])
		}
	}

	return len(wall) - max
}

func maxim(a, b int) int {
	if a > b {
		return a
	}
	return b
}
