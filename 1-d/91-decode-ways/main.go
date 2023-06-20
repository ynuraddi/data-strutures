package main

func main() {
	numDecodings("226")
}

// func numDecodings(s string) int {
// 	if s[0] == '0' {
// 		return 0
// 	}

// 	dp := make([]int, len(s)+1)
// 	str := []rune(s)
// 	dp[len(str)] = 1

// 	var dfs func(i int) int
// 	dfs = func(i int) int {
// 		if dp[i] != 0 {
// 			return dp[i]
// 		} else if str[i] == '0' {
// 			return 0
// 		}

// 		res := dfs(i + 1)
// 		if i+1 < len(str) && (str[i] == '1' || str[i] == '2' && str[i+1] <= '6') {
// 			res += dfs(i + 2)
// 		}
// 		dp[i] = res
// 		return res
// 	}
// 	sum := dfs(0)
// 	return sum
// }

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
		}

		if i+1 < len(s) && (s[i] == '1' || s[i] == '2' && (s[i+1] >= '0' && s[i+1] <= '6')) {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

// func numDecodings(s string) int {
// 	str := []rune(s)

// 	var dfs func(i int, s string) int
// 	dfs = func(i int, ss string) (sum int) {
// 		if i >= len(str) {
// 			fmt.Println(ss)
// 			return 1
// 		}

// 		if str[i] != '0' {
// 			sum += dfs(i+1, ss+" "+s[i:i+1])
// 			if i < len(str)-1 && (str[i] == '2' && str[i+1] <= '6' || str[i] == '1') {
// 				sum += dfs(i+2, ss+" "+s[i:i+2])
// 			}
// 		}
// 		return
// 	}
// 	sum := dfs(0, "")
// 	return sum
// }
