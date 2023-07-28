package main

func main() {
	minDistance("horse", "ros")
}

func minDistance(w1 string, w2 string) int {
	m := len(w1)
	n := len(w2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				ins := dp[i][j-1]
				del := dp[i-1][j]
				rep := dp[i-1][j-1]

				dp[i][j] = 1 + min(ins, min(rep, del))
			}
		}
	}

	return dp[m][n]
}

// func minDistance(w1 string, w2 string) int {
// 	dp := make([][]int, len(w1)+1)
// 	var a, b int = len(w1), len(w2)
// 	for i := range dp {
// 		dp[i] = make([]int, len(w2)+1)
// 		dp[i][len(w2)] = a
// 		a -= 1
// 	}
// 	for i := range dp[0] {
// 		dp[len(w1)][i] = b
// 		b -= 1
// 	}

// 	var dfs func(i, j int) int
// 	dfs = func(i, j int) int {
// 		if i == len(w1) || j == len(w2) {
// 			return dp[i][j]
// 		}

// 		if w1[i] == w2[j] {
// 			return dfs(i+1, j+1)
// 		}

// 		rep := dfs(i+1, j+1)
// 		del := dfs(i+1, j)
// 		ins := dfs(i, j+1)

// 		dp[i][j] = min(rep, min(del, ins)) + 1
// 		return dp[i][j]
// 	}

// 	return dfs(0, 0)
// }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
