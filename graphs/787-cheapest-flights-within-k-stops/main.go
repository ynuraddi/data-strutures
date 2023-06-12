package main

import (
	"fmt"
	"math"
)

func main() {
	// a := findCheapestPrice(3, [][]int{
	// 	{0, 1, 2},
	// 	{1, 2, 1},
	// 	{2, 0, 10},
	// }, 1, 2, 1)
	a := findCheapestPrice(4, [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}, 0, 3, 1)
	fmt.Println(a)
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([]int, n)

	for i := 0; i < n; i++ {
		cost[i] = math.MaxUint32
	}

	cost[src] = 0

	for i := 0; i <= k; i++ {
		temp := make([]int, n)
		copy(temp, cost)

		for j := 0; j < len(flights); j++ {
			from, to, price := flights[j][0], flights[j][1], flights[j][2]

			if cost[from] != math.MaxUint32 {
				temp[to] = min(temp[to], cost[from]+price)
			}
		}

		cost = temp
	}

	if cost[dst] == math.MaxUint32 {
		return -1
	}

	return cost[dst]
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
