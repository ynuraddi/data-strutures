package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

const input = `
8 10
1 2
1 3
1 4
4 3
3 2
2 4
1 8
5 6
7 6
5 7

`

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	users := make([][]int, n)
	usersF := make([]int, n)

	for i := 0; i < m; i++ {
		var user1, user2 int
		fmt.Fscan(in, &user1, &user2)
		user1, user2 = user1-1, user2-1

		// users[user1][user2] = -1
		// users[user2][user1] = -1
		users[user1] = append(users[user1], user2)
		users[user2] = append(users[user2], user1)

		usersF[user1] += 1
		usersF[user2] += 1
	}

	usersFF := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		if usersF[i] == 0 {
			continue
		}

		for j := 0; j < len(users[i]); j++ {
			JfI := users[i][j]
			mutualF := compareArrays(users[i], users[JfI], i, JfI)
			if len(mutualF) == 0 {
				continue
			}
			if usersFF[i] == nil {
				usersFF[i] = make(map[int]int, 5)
			}
			for _, u := range mutualF {
				usersFF[i][u] += 1
			}

		}
	}

	for _, uff := range usersFF {
		usersFriendsPrint(out, uff)
	}
}

func usersFriendsPrint(out io.Writer, uff map[int]int) {
	if uff == nil {
		fmt.Fprintln(out, 0)
		return
	}

	var max int
	for _, v := range uff {
		if max < v {
			max = v
		}
	}

	var result []int

	for k, v := range uff {
		if max != v {
			continue
		}
		result = append(result, k+1)
	}

	sort.Ints(result)

	for _, v := range result {
		fmt.Fprint(out, v, " ")
	}

	fmt.Fprintln(out)
}

func compareArrays(userF, FF []int, uid, fid int) []int {
	hash := make(map[int]bool, 5)

	for _, u := range userF {
		if u == fid {
			continue
		}
		hash[u] = true
	}

	result := make([]int, 0, 5)

	for _, u := range FF {
		if !hash[u] && u != uid {
			result = append(result, u)
		}
	}

	return result
}
