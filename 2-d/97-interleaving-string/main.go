package main

func main() {
	isInterleave("aabcc", "adbbca", "aadbbcabcac")
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	cache := make(map[[2]int]bool)

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i >= len(s1) && j >= len(s2) {
			return true
		}

		if val, ok := cache[[2]int{i, j}]; ok {
			return val
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}
		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}

		cache[[2]int{i, j}] = false
		return false
	}
	return dfs(0, 0)
}
