package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const input = `
4
A A A B
0 1 2 2 1 3 4 4 3 0

`

type Node struct {
	ID   int
	L    int // [a,b] = 0 , [a] = 1, [b] = 2
	Next map[int]*Node
}

func main() {
	in := strings.NewReader(input)
	// in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	workers := make([]*Node, 0, n+1)
	workers = append(workers, &Node{
		ID:   0,
		L:    0,
		Next: make(map[int]*Node, 1),
	})

	for i := 0; i < n; i++ {
		var lang string
		fmt.Fscan(in, &lang)

		var l int
		if lang == "A" {
			l = 1
		} else {
			l = 2
		}

		workers = append(workers, &Node{
			ID:   i + 1,
			L:    l,
			Next: make(map[int]*Node, 1),
		})
	}

	doc := make([]int, 0)
	var curr int
	fmt.Fscan(in, &curr)
	doc = append(doc, curr)

	nexts := make(map[int]int, n+1)
	nexts[0] = -1

	idx := 1
	for {
		fmt.Fscan(in, &curr)
		doc = append(doc, curr)
		nexts[curr] = idx
		idx++
		if curr == 0 {
			break
		}
	}

	var fill func(head, i, j int)
	fill = func(head, i, j int) {
		headNode := workers[doc[head]]
		for i := i + 1; i < j; i++ {
			headNode.Next[doc[i]] = workers[doc[i]]
			fill(i, i, nexts[doc[i]])
			i = nexts[doc[i]]
		}
	}

	fill(0, 0, nexts[0])

	result := make([]int, n+1)
	var dfs func(node *Node, a, b int)
	dfs = func(node *Node, a, b int) {
		if node.L == 1 {
			result[node.ID] = a
			a = -1
		} else {
			result[node.ID] = b
			b = -1
		}

		for _, next := range node.Next {
			dfs(next, a+1, b+1)
		}
	}

	dfs(workers[0], -1, -1)

	for i := 1; i < len(result); i++ {
		fmt.Fprint(out, result[i], " ")
	}

}
