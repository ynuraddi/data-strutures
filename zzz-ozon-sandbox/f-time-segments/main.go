package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// const input = `
// 6
// 1
// 02:46:00-03:14:59
// 2
// 23:59:59-23:59:59
// 00:00:00-23:59:58
// 2
// 23:59:58-23:59:59
// 00:00:00-23:59:58
// 2
// 23:59:59-23:59:58
// 00:00:00-23:59:57
// 6
// 17:53:39-20:20:02
// 10:39:17-11:00:52
// 08:42:47-09:02:14
// 09:44:26-10:21:41
// 00:46:17-02:07:19
// 22:42:50-23:17:46
// 1
// 24:00:00-23:59:59

// `

const (
	YES = "YES"
	NO  = "NO"
)

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)

		timezone := make([]bool, 60*60*24)
		var incorrect bool

		for i := 0; i < t; i++ {
			var time string
			fmt.Fscan(in, &time)

			if incorrect {
				continue
			}

			var h1, m1, s1 int
			var h2, m2, s2 int

			h1, _ = strconv.Atoi(time[0:2])
			m1, _ = strconv.Atoi(time[3:5])
			s1, _ = strconv.Atoi(time[6:8])
			h2, _ = strconv.Atoi(time[9:11])
			m2, _ = strconv.Atoi(time[12:14])
			s2, _ = strconv.Atoi(time[15:])

			if h1 < 0 || h1 > 23 || h2 < 0 || h2 > 23 {
				incorrect = true
				continue
			} else if m1 < 0 || m1 > 59 || m2 < 0 || m2 > 59 {
				incorrect = true
				continue
			} else if s1 < 0 || s1 > 59 || s2 < 0 || s2 > 59 {
				incorrect = true
				continue
			}

			t1 := s1 + 60*m1 + 3600*h1
			t2 := s2 + 60*m2 + 3600*h2

			if t1 > t2 {
				incorrect = true
				continue
			}

			for i := t1; i <= t2; i++ {
				if !timezone[i] {
					timezone[i] = true
				} else {
					incorrect = true
					continue
				}
			}

			// fmt.Println(h1, m1, s1, h2, m2, s2)
		}

		if incorrect {
			fmt.Fprintln(out, NO)
		} else {
			fmt.Fprintln(out, YES)
		}

	}
}
