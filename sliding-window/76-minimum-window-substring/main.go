package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	sLength := len(s)
	tLength := len(t)
	set := make(map[byte][]int)
	for i := 0; i < len(t); i++ {
		if _, ok := set[t[i]]; !ok {
			set[t[i]] = append(set[t[i]], 1, 0)
		} else {
			set[t[i]][0]++
		}
	}

	left := 0
	matches := 0
	var result string

	for right := 0; right < sLength; right++ {
		if _, ok := set[s[right]]; ok {
			set[s[right]][1]++

			if set[s[right]][0] >= set[s[right]][1] {
				matches++
			}
		}

		for matches == tLength {
			if right-left+1 < len(result) || len(result) == 0 {
				result = s[left : right+1]
			}

			if _, ok := set[s[left]]; ok {
				set[s[left]][1]--
				if set[s[left]][0] > set[s[left]][1] {
					matches--
				}
			}

			left++
		}
	}

	return result
}

// for r < len(s) && hash != [26]int{} {
// 	if s[r] == s[l-1] {
// 		hash[s[r]%26]--
// 		break
// 	}
// 	if hashT[s[r]%26] != 0 {
// 		hash[s[r]%26]--
// 	}
// 	r++
// }

// for l < len(s) && hash != [26]int{} {
// 	if hash[s[l]%26] < 0 {
// 		hash[s[l]%26]++
// 		break
// 	}
// 	if hashT[s[l]%26] != 0 {
// 		hashT[s[l]%26]--
// 	}
// 	l++
// }
