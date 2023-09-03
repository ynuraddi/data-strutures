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

	var n, money int
	fmt.Fscan(in, &n, &money)

	var max int

	for i := 0; i < n; i++ {
		var price int
		fmt.Fscan(in, &price)

		if price <= money {
			max = Max(max, price)
		}
	}

	fmt.Fprintln(out, max)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
