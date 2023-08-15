package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const input = `
20
sss
ddvlwswlsd
swvw
dl
llwllvw
swdsdwsvd
vwvsv
wslswds
vwswll
wwvdsslwd
wsvvlsw
lvvlvsdss
ldwlvsd
v
wvldwd
ldlddvddvv
vsswwll
svwl
sldwdsvsv
svdw
20
vwlldlswl
slvlvdldw
vwls
dl
lvs
ldvsl
ddlsw
d
svsvvwsdwl
vlssldswlw
lvlsdl
s
vlwlssld
ldvl
dvvvwsssww
vdlvldw
wwv
dlwlwwsvdl
swsw
dwldddwwl

`

type TrieNode struct {
	tail bool
	head *TrieNode

	next map[byte]*TrieNode
}

func (t *TrieNode) Insert(word string) {
	this := t

	for i := len(word) - 1; i >= 0; i-- {
		if _, ok := this.next[word[i]]; !ok {
			this.next[word[i]] = &TrieNode{
				next: make(map[byte]*TrieNode, 26),
			}
		}
		this = this.next[word[i]]
	}
	this.tail = true
}

func (t *TrieNode) FindWord(word string) (string, int) {
	this := t
	currWord, currLen := "", 0

	for i := len(word) - 1; i >= 0; i-- {
		if _, ok := this.next[word[i]]; !ok {
			break
		}

		currWord = string(word[i]) + currWord
		currLen += 1
	}

	if currWord == word && len(this.next) == 0 {

		this = this.head
		// dontTouch := currWord[len(currWord)-1]
		currLen--

	}

	return this.endW() + currWord, currLen
}

func (t *TrieNode) endW() string {
	this := t

	for let, node := range this.next {
		return node.endW() + string(let)
	}
	return ""
}

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	in := strings.NewReader(input)
	// in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	root := &TrieNode{
		next: make(map[byte]*TrieNode, 26),
	}

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)

		root.Insert(word)
	}

	var m int
	fmt.Fscan(in, &m)

	for i := 0; i < m; i++ {
		var word string
		fmt.Fscan(in, &word)

		w, _ := root.FindWord(word)
		fmt.Fprintln(out, w)

	}
}
