package main

func main() {
	removeStars("leet**cod*e*")
}

func removeStars(s string) string {
	stack := make([]byte, len(s))

	var idx int

	for i := range s {
		if s[i] == '*' {
			idx -= 1
		} else {
			stack[idx] = s[i]
			idx += 1
		}
	}

	return string(stack[:idx])
}
