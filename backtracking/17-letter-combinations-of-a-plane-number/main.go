package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) (res []string) {
	if len(digits) < 1 {
		return nil
	}

	letters := map[byte]string{
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	tmp := ""

	var bfs func(idx int)
	bfs = func(idx int) {
		if idx == len(digits) {
			res = append(res, tmp)
			return
		}

		for _, v := range letters[digits[idx]] {
			tmp += string(v)
			bfs(idx + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	bfs(0)
	return res
}
