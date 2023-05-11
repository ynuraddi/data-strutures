package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkInclusion("ab", "oba"))
}

func checkInclusion(s1 string, s2 string) bool {
	l, count := 0, [26]int{}
	for i := range s1 {
		count[s1[i]-97]++
	}

	for r := range s2 {
		count[s2[r]-97]--
		if count == [26]int{} {
			return true
		}

		if r+1 >= len(s1) {
			count[s2[l]-97]++
			l++
		}
	}
	return false
}

// func checkInclusion(s1 string, s2 string) bool {
// 	if len(s1) > len(s2) {
// 		return false
// 	}

// 	hash1 := make([]int, 26)
// 	hash2 := make([]int, 26)

// 	for i := range s1 {
// 		hash1[s1[i]%26]++
// 	}

// 	for i := 0; i < len(s1)-1; i++ {
// 		hash2[s2[i]%26]++
// 	}

// 	for i, j := 0, len(s1)-1; j < len(s2); i, j = i+1, j+1 {
// 		indexLetter := s2[j] % 26
// 		hash2[indexLetter]++
// 		if hash1[indexLetter] == hash2[indexLetter] {
// 			if reflect.DeepEqual(hash1, hash2) {
// 				return true
// 			}
// 		}

// 		hash2[s2[i]%26]--
// 	}

// 	return false
// }
