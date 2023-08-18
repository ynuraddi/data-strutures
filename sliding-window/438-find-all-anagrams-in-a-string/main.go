package main

func main() {
	findAnagrams("cbaebabacd", "abc")
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	results := make([]int, 0, 10)

	mask := [26]int{}
	curr := [26]int{}

	for i := range p {
		mask[p[i]%26] += 1
		curr[s[i]%26] += 1
	}

	if mask == curr {
		results = append(results, 0)
	}

	for i, j := 0, len(p); i < len(s) && j < len(s); i, j = i+1, j+1 {
		curr[s[j]%26] += 1
		curr[s[i]%26] -= 1

		if mask == curr {
			results = append(results, i+1)
		}
	}

	return results
}
