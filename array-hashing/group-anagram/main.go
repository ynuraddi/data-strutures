package main

func main() {
}

type CharCount [26]int

func groupAnagrams(strs []string) [][]string {
	hash := make(map[CharCount][]string)

	for _, word := range strs {
		charCount := CharCount{}
		for i := range word {
			charCount[word[i]-'a']++
		}

		hash[charCount] = append(hash[charCount], word)
	}

	var result [][]string
	for _, v := range hash {
		result = append(result, v)
	}

	return result
}
