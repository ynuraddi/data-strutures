package main

import (
	"bufio"
	"fmt"
	"os"
)

// const input = `
// 7 13
// 2 3 1
// 3 3
// 1 2 4
// 2 1 1
// 3 4
// 2 3 4
// 1 3 4
// 3 4
// 1 7 3
// 1 1 3
// 3 7
// 3 1
// 2 7 4
// `

func main() {
	in := bufio.NewReader(os.Stdin)
	// in := strings.NewReader(input)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	hist := make([][]int, n+1)
	band := make(map[int][]int)
	for i := 1; i <= n; i++ {
		hist[i] = append(hist[i], i)
		band[i] = append(band[i], i)
	}

	bandCounter := n + 1
	for i := 0; i < m; i++ {
		var cmd int
		fmt.Fscan(in, &cmd)

		if cmd == 1 {
			var h1, h2 int
			fmt.Fscan(in, &h1, &h2)

			if last(hist[h1]) == last(hist[h2]) {
				continue
			}

			band[bandCounter] = append(band[last(hist[h1])], band[last(hist[h2])]...)
			for _, h := range band[bandCounter] {
				delete(band, last(hist[h]))
				hist[h] = append(hist[h], bandCounter)
			}
			bandCounter += 1
		} else if cmd == 2 {
			var h1, h2 int
			fmt.Fscan(in, &h1, &h2)

			if last(hist[h1]) == last(hist[h2]) {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		} else if cmd == 3 {
			var h int
			fmt.Fscan(in, &h)

			fmt.Fprintln(out, len(hist[h]))
		}
	}
}

func last(arr []int) int {
	return arr[len(arr)-1]
}
