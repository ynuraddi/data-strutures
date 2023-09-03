package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const input = `
2 2 3
a b
1 1
2 2
a < 3
b > 1
b < 3

`

func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	// 0 - index, 1 - min, 2 - max
	columnsName := make(map[string][]int)

	for i := 0; i < m; i++ {
		var name string
		fmt.Fscan(in, &name)
		columnsName[name] = []int{i, math.MinInt, math.MaxInt}
	}

	plain := make([][]int, n)
	for i := range plain {
		plain[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var n int
			fmt.Fscan(in, &n)

			plain[i][j] = n
		}
	}

	for i := 0; i < q; i++ {
		var column, con string
		var val int
		fmt.Fscan(in, &column, &con, &val)

		if con == "<" {
			columnsName[column][2] = Min(columnsName[column][2], val)
		} else {
			columnsName[column][1] = Max(columnsName[column][1], val)
		}
	}

	maxmin := make(map[int][2]int, len(columnsName))
	for _, v := range columnsName {
		maxmin[v[0]] = [2]int{v[1], v[2]}
	}

	sum := 0
	for i := range plain {
		localsum := 0
		for j, v := range plain[i] {
			if maxmin[j][0] < v && v < maxmin[j][1] {
				localsum += v
			} else {
				localsum = 0
				break
			}
		}
		sum += localsum
	}

	fmt.Fprintln(out, sum)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
