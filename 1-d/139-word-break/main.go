package main

func main() {
	wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s); i >= 0; i-- {
		for _, word := range wordDict {
			if dp[i] && i >= len(word) && s[i-len(word):i] == word {
				dp[i-len(word)] = true
			}
		}
	}

	return dp[0]
}
