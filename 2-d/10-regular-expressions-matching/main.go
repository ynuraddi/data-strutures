package main

func main() {
	isMatch("afjvnaaaaa", ".*aaa")
}

// HE-HE
// func isMatch(s string, p string) bool {
//     pattern := regexp.MustCompile(p)
//     return s == pattern.FindString(s)
// }

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			} else if i == 0 {
				dp[i][j] = ((j-1)%2 == 1 && p[j-1] == '*' && dp[i][j-2])
				continue
			} else if j == 0 {
				dp[i][j] = false
				continue
			}

			if p[j-1] != '*' {
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '.') && dp[i-1][j-1]
			} else {
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i-1][j]
				}
				if dp[i][j-2] {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}
