package main

import (
	"bufio"
	"fmt"
	"os"
)

const input = `
3
1
5
2
1 2
3
5 1 2

`

type Node struct {
	countMass int
	Next      map[int]*Node
}

func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	sum := 0

	root := &Node{
		countMass: 0,
		Next:      make(map[int]*Node),
	}

	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)

		this := root

		for i := 0; i < k; i++ {
			var num int
			fmt.Fscan(in, &num)

			if _, ok := this.Next[num]; !ok {
				this.Next[num] = &Node{
					countMass: 0,
					Next:      make(map[int]*Node),
				}
			}
			this = this.Next[num]
			sum += this.countMass
			this.countMass += 1
		}
	}

	fmt.Fprintln(out, sum)
}
