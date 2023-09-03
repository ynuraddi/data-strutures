package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const input = `
11 3
1 2 5
`

func main() {
	// in := bufio.NewReader(os.Stdin)
	in := strings.NewReader(input)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	nomins := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var nom int
		fmt.Fscan(in, &nom)

		nomins = append(nomins, nom)
	}

	var result []int
	var finish bool

	var bfs func(i int, sum int, curr []int)
	bfs = func(i, sum int, curr []int) {
		if sum == n && !finish {
			result = append([]int{}, curr...)
			finish = true
			return
		}

		if sum > n || finish || i >= len(nomins) {
			return
		}

		bfs(i+1, sum, curr)

		curr = append(curr, nomins[i])
		bfs(i+1, sum+nomins[i], curr)

		curr = append(curr, nomins[i])
		bfs(i+1, sum+nomins[i]*2, curr)
	}
	bfs(0, 0, result)

	if len(result) == 0 {
		fmt.Fprintln(out, -1)
		return
	}

	fmt.Fprintln(out, len(result))
	for _, v := range result {
		fmt.Fprint(out, v, " ")
	}
}
