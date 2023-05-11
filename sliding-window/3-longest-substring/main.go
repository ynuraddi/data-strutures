package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]bool, 26)
	result := 0
	tmp := 0

	for i, j := 0, 0; j < len(s); j++ {
		if !hash[s[j]] {
			hash[s[j]] = true
			tmp++
		} else {
			if tmp > result {
				result = tmp
			}

			if s[i] == s[j] {
				i++
			} else {
				for s[i] != s[j] {
					hash[s[i]] = false
					tmp--
					i++
				}
				i++
			}
		}
	}

	if tmp > result {
		return tmp
	}

	return result
}
