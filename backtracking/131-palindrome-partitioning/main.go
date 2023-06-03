package main

func main() {
}

func partition(s string) (res [][]string) {
	curr := make([]string, 0)
	var bfs func(idx int)
	bfs = func(idx int) {
		if idx == len(s) {
			res = append(res, append([]string{}, curr...))
		}
		for i := idx; i < len(s); i++ {
			if isPalidrome(s[idx : i+1]) {
				curr = append(curr, s[idx:i+1])
				bfs(i + 1)
				curr = curr[:len(curr)-1]
			}
		}
	}
	bfs(0)
	return
}

func isPalidrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
