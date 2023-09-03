package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	s1 := make([]int, n)
	hash1 := make(map[int]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s1[i])
		hash1[s1[i]] += 1
	}

	s2 := make([]int, n)
	hash2 := make(map[int]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s2[i])
		hash2[s2[i]] += 1
	}

	var start, end int = -1, n
	for i := 0; i < n && s1[i] == s2[i]; i++ {
		start = i
	}
	for i := n - 1; i >= 0 && s1[i] == s2[i]; i-- {
		end = i
	}

	start += 1
	for start < end {
		if start < end-1 && s2[start] > s2[start+1] {
			fmt.Fprintln(out, "NO")
			return
		}

		if hash1[s1[start]] != hash2[s1[start]] {
			fmt.Fprintln(out, "NO")
			return
		}
		start += 1
	}

	fmt.Fprintln(out, "YES")
}
