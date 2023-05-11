package main

import (
	"fmt"
)

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	hashOfCount := make([]int, 26)
	result := 0
	temp := 0

	for i, j := 0, 0; j < len(s); {
		indexJ := s[j] % 26
		indexI := s[i] % 26
		hashOfCount[indexJ]++
		temp++

		length := j - i + 1
		power := length - max(hashOfCount)

		if power <= k {
			j++
		} else {
			hashOfCount[indexI]--
			temp--

			j++
			i++
		}

		if power == k && result < temp {
			result = temp
		}
	}

	if temp > result {
		return temp
	}

	return result
}

func max(a []int) (max int) {
	for _, v := range a {
		if max < v {
			max = v
		}
	}
	return
}
