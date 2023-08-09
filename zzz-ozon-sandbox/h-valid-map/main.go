package main

import (
	"bufio"
	"fmt"
	"os"
)

// const input = `
// 1
// 4 4
// R.R.
// .Y.R
// Y.K.
// .B.R

// `

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)

		mapa := make([][]rune, x)
		for i := range mapa {
			mapa[i] = make([]rune, y)
		}

		for i := range mapa {
			var line string
			fmt.Fscan(in, &line)

			v := []rune(line)
			for j := range mapa[i] {
				mapa[i][j] = v[j]
			}
		}

		usedColor := make([]bool, 150)

		var bfs func(i, j int, target rune)
		bfs = func(i, j int, target rune) {
			if i < 0 || i >= x ||
				j < 0 || j >= y ||
				mapa[i][j] != target ||
				mapa[i][j] == '.' {
				return
			}

			mapa[i][j] = '.'

			bfs(i-1, j+1, target)
			bfs(i-1, j-1, target)
			bfs(i, j-2, target)
			bfs(i, j+2, target)
			bfs(i+1, j+1, target)
			bfs(i+1, j-1, target)
		}

		var incorrect bool

		for i := range mapa {
			var pov int
			if i%2 != 0 {
				pov += 1
			}

			for j := pov; j < y; j += 2 {
				if mapa[i][j] == '.' {
					continue
				}
				if usedColor[mapa[i][j]] {
					incorrect = true
					break
				}
				usedColor[mapa[i][j]] = true
				bfs(i, j, mapa[i][j])
			}
		}

		if incorrect {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
