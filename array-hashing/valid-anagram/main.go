package main

func main() {
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashS := make([]int, 26)
	hashT := make([]int, 26)

	for _, v := range s {
		hashS[v%26]++
	}

	for _, v := range t {
		hashT[v%26]++
	}

	for i := range hashS {
		if hashS[i] != hashT[i] {
			return false
		}
	}

	return true
}
