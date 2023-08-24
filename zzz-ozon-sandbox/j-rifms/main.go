package main

import (
	"bufio"
	"fmt"
	"os"
)

const input = `
10
iiii
iiiii
iiiiiiii
iiiiiiiii
iii
ii
i
iiiiiii
iiiiii
iiiiiiiiii
20
iiiiii
iiii
iiiiiiiiii
iiiiiiii
ii
iiiiiii
iiiiiiii
iiiiiiii
i
i
iiiiii
i
i
iiiii
iii
iii
iiiiiiiiii
iiii
iiiiiiiiii
iii




`

type Node struct {
	tail bool

	next map[string]*Node
}

func (t *Node) Insert(word string) {
	this := t

	for i := len(word) - 1; i >= 0; i-- {
		curr := string(word[i])

		if _, ok := this.next[curr]; !ok {
			this.next[curr] = &Node{
				next: make(map[string]*Node, 26),
			}
		}

		this = this.next[curr]
	}

	this.tail = true
}

func (t *Node) Find(word string) string {
	this := t

	var result string

	history := make([]*Node, 0, len(word))
	historyRes := make([]string, 0, len(word))

	for i := len(word) - 1; i >= 0; i-- {
		curr := string(word[i])

		if _, ok := this.next[curr]; !ok {
			break
		}

		history = append(history, this)
		historyRes = append(historyRes, result)

		this = this.next[curr]
		result = curr + result
	}

	if !this.tail {
		return this.endWord() + result
	}

	if result != word {
		return result
	}

	if len(this.next) != 0 {
		return this.endWord() + result
	}

	for i := len(history) - 1; i >= 0; i-- {
		if history[i].tail {
			return historyRes[i]
		}

		if len(history[i].next) == 1 {
			continue
		}

		curr := string(word[len(word)-1-i])
		for let, next := range history[i].next {
			if let != curr {
				return next.endWord() + let + historyRes[i]
			}
		}

		break
	}

	return result
}

func (t *Node) endWord() string {
	this := t

	for let, next := range this.next {
		res := next.endWord()
		return res + let
	}

	return ""
}

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	root := &Node{
		next: make(map[string]*Node, 26),
	}

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)

		root.Insert(word)
	}

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)

		result := root.Find(word)

		fmt.Fprintln(out, result)
	}
}
