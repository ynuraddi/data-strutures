package main

func main() {
	// wd := Constructor()

	// wd.AddWord("at")
	// wd.AddWord("and")
	// wd.AddWord("an")
	// wd.AddWord("add")

	// b := wd.Search("a")
	// b = wd.Search(".at")

	// wd.AddWord("bat")

	// b = wd.Search(".at")
	// b = wd.Search("an.")
	// b = wd.Search("a.d.")
	// b = wd.Search("b.")
	// b = wd.Search("a.d")
	// b = wd.Search(".")

	// fmt.Println(b)
}

// type TrieNode struct {
// 	children [25]*TrieNode
// 	word     bool
// }

// type WordDictionary struct {
// 	root *TrieNode
// }

// func Constructor() WordDictionary {
// 	return WordDictionary{root: &TrieNode{children: [25]*TrieNode{}}}
// }

// func (this *WordDictionary) AddWord(word string) {
// 	cur := this.root
// 	for c := 0; c < len(word); c++ {
// 		if cur.children[word[c]-'a'] == nil {
// 			cur.children[word[c]-'a'] = &TrieNode{children: [25]*TrieNode{}}
// 		}
// 		cur = cur.children[word[c]-'a']
// 	}
// 	cur.word = true
// }

// func (this *WordDictionary) Search(word string) bool {
// 	cur := this.root

// 	for c := 0; c < len(word); c++ {
// 		if word[c] == '.' {

// 		}
// 	}

// 	return cur.word
// }
