package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// const input = `
// 5 6
// 1 2 8
// 2 3 6
// 2 3 2
// 3 1 4
// 5 4 1
// 4 5 8
// `

type City struct {
	ID    int
	Roads map[int]int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	// in := strings.NewReader(input)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	cities := make(map[int]*City, n)
	for i := 0; i < m; i++ {
		var c1, c2, l int
		fmt.Fscan(in, &c1, &c2, &l)

		if _, ok := cities[c1]; !ok {
			cities[c1] = &City{
				ID:    c1,
				Roads: make(map[int]int),
			}
		}
		if _, ok := cities[c2]; !ok {
			cities[c2] = &City{
				ID:    c2,
				Roads: make(map[int]int),
			}
		}
		cities[c1].Roads[c2] = Max(cities[c1].Roads[c2], l)
		cities[c2].Roads[c1] = Max(cities[c2].Roads[c1], l)
	}

	cityCost := make(map[int]int)
	dfs := func(cid, minnum int) {
		curr := cities[cid]
		for ncid, road := range curr.Roads {
			cityCost[ncid] = Max(cityCost[ncid], road)
		}
	}

	for i := range cities {
		dfs(i, 0)
	}

	result := math.MaxInt
	for _, v := range cityCost {
		result = Min(v, result)
	}

	fmt.Fprintln(out, result-1)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
