package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("123abobbbbbbbbbbbbaxyz"))
}

func longestPalindrome(s string) string {
	plain := make([][]int, len(s))
	for i := range plain {
		plain[i] = make([]int, len(s))
	}

	result := ""
	maxim := 0

	for j := len(s) - 1; j >= 0; j-- {
		jpos := len(s) - 1 - j
		for i := 0; i < len(s); i++ {
			if s[i] != s[j] {
				continue
			}

			if i == 0 || jpos == 0 {
				plain[jpos][i] = 1
			} else {
				plain[jpos][i] = plain[jpos-1][i-1] + 1
			}

			if plain[jpos][i] > maxim {
				maxim = plain[jpos][i]
				result = s[i-maxim+1 : i+1]
			}
		}
	}

	return result
}
