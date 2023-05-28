package main

func main() {
}

type Trie struct {
	children  [26]*Trie
	endOfWord bool
	root      *Trie
}

func Constructor() Trie {
	return Trie{
		children:  [26]*Trie{},
		endOfWord: false,
		root:      &Trie{},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root

	for _, r := range word {
		if curr.children[r-'a'] == nil {
			curr.children[r-'a'] = &Trie{}
		}
		curr = curr.children[r-'a']
	}

	curr.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root

	for _, r := range word {
		if curr.children[r-'a'] == nil {
			return false
		}
		curr = curr.children[r-'a']
	}

	return curr.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root

	for _, r := range prefix {
		if curr.children[r-'a'] == nil {
			return false
		}
		curr = curr.children[r-'a']
	}

	return true
}
