package main

import (
	"bufio"
	"fmt"
	"os"
)

// const input = `
// 6
// 12
// 2 2 2 2 2 2 2 3 3 3 3 3
// 12
// 2 3 2 3 2 2 3 2 3 2 2 3
// 1
// 10000
// 9
// 1 2 3 1 2 3 1 2 3
// 6
// 10000 10000 10000 10000 10000 10000
// 6
// 300 100 200 300 200 300
// `

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var size int
	var arr map[int]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &size)

		arr = make(map[int]int, size)
		for j := 0; j < size; j++ {
			var val int
			fmt.Fscan(in, &val)

			arr[val] += 1
		}
		// fmt.Println(size, arr)

		fmt.Println(findCost(arr))
	}
}

func findCost(a map[int]int) int {
	var sum int

	for key, val := range a {
		sum += (val - val/3) * key
	}

	return sum
}
