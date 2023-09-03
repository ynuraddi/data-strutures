package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	if len(s) < len("sheriff") {
		fmt.Fprintln(out, 0)
		return
	}

	word := "sherif"

	hash := make(map[byte]int, len(word))
	for i := range s {
		if s[i] == 's' || s[i] == 'h' || s[i] == 'e' || s[i] == 'r' || s[i] == 'i' || s[i] == 'f' {
			hash[s[i]] += 1
		}
	}

	min := math.MaxInt
	for i := range word {
		if word[i] == 'f' {
			min = Min(min, hash[word[i]]/2)
		} else {
			min = Min(min, hash[word[i]])
		}
	}

	fmt.Fprintln(out, min)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
