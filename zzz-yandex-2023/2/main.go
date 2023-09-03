package main

import (
	"bufio"
	"fmt"
	"os"
)

// const input = `
// 3 3 20
// 1 6 7
// 2 4 5
// 1 A 2
// 1 B 1
// 1 B 8
// 1 B 5
// 1 A 3
// 1 A 2
// 1 B 10
// 1 A 9
// 1 A 8
// 1 B 7
// -1 A 1
// -1 B 5
// -1 B 5
// -1 B 4
// -1 A 6
// -1 A 8
// -1 A 2
// -1 B 8
// -1 B 10
// -1 A 2

// `

func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	collBoth := make(map[int]int, n)

	collA := make(map[int]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		collBoth[k] += 1
		collA[k] += 1
	}

	collB := make(map[int]int, m)
	for i := 0; i < m; i++ {
		var k int
		fmt.Fscan(in, &k)
		collBoth[k] += 1
		collB[k] += 1
	}

	diff := 0
	for k := range collBoth {
		diff += abs(collA[k] - collB[k])
	}

	for i := 0; i < q; i++ {
		var pl string
		var typ, card int
		fmt.Fscan(in, &typ, &pl, &card)

		if typ == 1 {
			if pl == "A" {
				if collB[card] > collA[card] {
					diff -= 1
				} else if collB[card] <= collA[card] {
					diff += 1
				}
				collA[card] += 1
			} else {
				if collA[card] > collB[card] {
					diff -= 1
				} else if collA[card] <= collB[card] {
					diff += 1
				}
				collB[card] += 1
			}
		} else {
			if pl == "A" {
				if collB[card] >= collA[card] {
					diff += 1
				} else if collB[card] < collA[card] {
					diff -= 1
				}
				collA[card] -= 1
			} else {
				if collA[card] >= collB[card] {
					diff += 1
				} else if collA[card] < collB[card] {
					diff -= 1
				}
				collB[card] -= 1
			}
		}

		fmt.Fprint(out, diff, " ")
	}

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
