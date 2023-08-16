package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const input = `
*****.**
***.*.**
**..****
**..*...
........
****.***
...*.***
******..
1
3
***
..*
..*
`

func NewFigure(s ...string) [][]byte {
	n := len(s)

	f := make([][]byte, n)

	for _, v := range s {
		var tmp []byte
		for j := range v {
			tmp = append(tmp, v[j])
		}
		f = append(f, tmp)
	}

	return f
}

type plain struct {
	m     [8][8]byte
	count int
}

func (p *plain) IsInsert(s []string, i, j int) bool {
	for ; i < 8-len(s); i++ {
		for ; j < 8-len(s); j++ {
			if p.m[i][j] == '*' && s[i][j] == '*' {
				return false
			}
		}
	}

	return false
}

func (p *plain) Insert(s []string, i, j int) {
	for ; i < 8-len(s); i++ {
		for ; j < 8-len(s); j++ {
			if p.m[i][j] == '.' && s[i][j] == '*' {
				p.m[i][j] = '*'
			}
		}
	}
}

func (p *plain) Delete(s []string, i, j int) {
	for ; i < 8-len(s); i++ {
		for ; j < 8-len(s); j++ {
			if p.m[i][j] == '*' && s[i][j] == '*' {
				p.m[i][j] = '.'
			}
		}
	}
}

func (p *plain) Check(i int) (is, js [8]int) {
	for i := range p.m {
		for j := range p.m[i] {
			if p.m[i][j] == '*' {
				is[i] += 1
				js[j] += 1
			}
		}
	}
	return
}

func main() {
	in := strings.NewReader(input)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	plain := plain{
		m: [8][8]byte{},
	}

	for i := 0; i < 8; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < 8; j++ {
			if s[j] == '*' {
				plain.count++
			}
			plain.m[i][j] = s[j]
		}
	}

	for i := 0; i < 8; i++ {
		fmt.Fprintln(out, plain.m[i])
	}
	fmt.Fprintln(out, plain.count)

	var n int
	fmt.Fscan(in, &n)

	figures := make([][]string, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(in, &m)
		for j := 0; j < m; j++ {
			var tmp string
			fmt.Fscan(in, &tmp)
			figures[i] = append(figures[i], tmp)
		}
	}

	fmt.Fprintln(out, figures)

	var bfs func(f [][]string) int
	bfs = func(f [][]string) int {
		if len(f) == 0 {
			return plain.count
		}

		for fidx, fig := range f {
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					if plain.IsInsert(fig, i, j) {
						plain.Insert(fig, i, j)

						plain.Delete(fig, i, j)
					}
				}
			}
		}
	}
}

func min(a ...int) int {
	min := math.MaxInt
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
