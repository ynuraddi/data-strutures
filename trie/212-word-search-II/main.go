package main

import "fmt"

func main() {
	a := findWords([][]byte{
		{'a', 'b', 'e'},
		{'b', 'c', 'd'},
	}, []string{"abcdeb"})
	fmt.Println(a)

	findWords([][]byte{
		{'a', 'a'},
		{'a', 'a'},
	}, []string{"aaaaa"})
	findWords([][]byte{
		{'a', 'b', 'c'},
		{'a', 'e', 'd'},
		{'a', 'f', 'g'},
	}, []string{"eaabcdgfa"})
}

type TrieNode struct {
	children   map[byte]*TrieNode
	endOfWord  bool
	references int
}

func (t *TrieNode) addWord(word string) {
	cur := t
	cur.references++

	for i := range word {
		if _, ok := cur.children[word[i]]; !ok {
			cur.children[word[i]] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		cur = cur.children[word[i]]
		cur.references++
	}
	cur.endOfWord = true
}

func (t *TrieNode) removeWord(word string) {
	cur := t
	cur.references--

	for i := range word {
		if _, ok := cur.children[word[i]]; ok {
			cur = cur.children[word[i]]
			cur.references--
		}
	}
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, w := range words {
		root.addWord(w)
	}

	ROWS, COLS := len(board), len(board[0])
	res := make(map[string]bool)
	visit := make(map[int]bool)

	var dfs func(x, y int, t *TrieNode, word string)
	dfs = func(x, y int, t *TrieNode, word string) {
		if x < 0 || x >= ROWS ||
			y < 0 || y >= COLS ||
			t.children[board[x][y]] == nil ||
			t.children[board[x][y]].references < 1 ||
			visit[x*COLS+y] {
			return
		}

		visit[x*COLS+y] = true
		t = t.children[board[x][y]]
		word += string(board[x][y])
		if t.endOfWord {
			t.endOfWord = false
			res[word] = true
			root.removeWord(word)
		}

		dfs(x+1, y, t, word)
		dfs(x-1, y, t, word)
		dfs(x, y+1, t, word)
		dfs(x, y-1, t, word)
		visit[x*COLS+y] = false
	}

	for x := range board {
		for y := range board[x] {
			dfs(x, y, root, "")
		}
	}

	result := make([]string, 0, len(words))
	for w := range res {
		result = append(result, w)
	}
	return result
}
