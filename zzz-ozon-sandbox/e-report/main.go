package main

import (
	"bufio"
	"fmt"
	"os"
)

// const input = `
// 4
// 8
// 3 8 8 8 3 8 8 8
// 9
// 1 2 3 4 5 4 3 2 1
// 7
// 7 1 1 1 1 2 2
// 6
// 3 3 4 4 5 5

// `

const (
	YES = "YES"
	NO  = "NO"
)

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var countOfInstructions int
		fmt.Fscan(in, &countOfInstructions)

		startedAssignments := make(map[int]bool, countOfInstructions)

		var prev int
		var isBad bool
		for i := 0; i < countOfInstructions; i++ {
			var curr int
			fmt.Fscan(in, &curr)

			if isBad {
				continue
			}

			if startedAssignments[curr] {
				isBad = true
			}

			if i > 0 && prev != curr {
				startedAssignments[prev] = true
			}
			prev = curr
		}
		if isBad {
			fmt.Fprintln(out, NO)
		} else {
			fmt.Fprintln(out, YES)
		}
	}
}
