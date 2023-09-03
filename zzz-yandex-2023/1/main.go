package main

import (
	"bufio"
	"fmt"
	"os"
)

// const input = `
// 1001 5 20 14 15 16
// 9009 9 11 12 21 11
// `

func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var y1, m1, d1, h1, min1, s1 int
	var y2, m2, d2, h2, min2, s2 int

	fmt.Fscan(in, &y1, &m1, &d1, &h1, &min1, &s1)
	fmt.Fscan(in, &y2, &m2, &d2, &h2, &min2, &s2)

	daymonths := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dm1 := 0
	dm2 := 0
	for i, v := range daymonths {
		if i+1 < m1 {
			dm1 += v
		}
		if i+1 < m2 {
			dm2 += v
		}
	}

	days1 := y1*365 + dm1 + d1
	days2 := y2*365 + dm2 + d2

	diffdays := days2 - days1

	secs1 := s1 + min1*60 + h1*3600
	secs2 := s2 + min2*60 + h2*3600

	if secs2 < secs1 {
		diffdays -= 1
		secs2 += 86400
	}

	diffsecs := secs2 - secs1

	fmt.Fprintln(out, diffdays, diffsecs)
}
