package main

func main() {
	longestPalindrome("babad")
}

func longestPalindrome(s string) string {
	str := []rune(s)
	start, end := 0, 0
	for i := 0; i < len(str); i++ {
		len1 := Centre(str, i, i)
		len2 := Centre(str, i, i+1)
		if (len1 > end-start) || (len2 > end-start) {
			if len1 > len2 {
				start = i - (len1-1)/2
				end = i + len1/2
			} else {
				start = i - (len2-1)/2
				end = i + len2/2
			}
		}
	}
	return string(str[start : end+1])
}

func Centre(s []rune, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
