package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	s = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
