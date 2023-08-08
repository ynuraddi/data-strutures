package main

import (
	"bufio"
	"fmt"
	"os"
)

// const input = `
// 8 6
// 4 3
// 3 1
// 1 2
// 2 4
// 2 5
// 6 8

// `

// dinamic caching
type User struct {
	ID      int
	Friends []*User

	countOfFriends int

	maxFriend         int
	maxFriendWithUser []int
}

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	users := make([]*User, n)
	for i := range users {
		users[i] = &User{
			ID:      i,
			Friends: make([]*User, n),

			maxFriendWithUser: make([]int, n),
		}
	}

	for i := 0; i < m; i++ {
		var id1, id2 int
		fmt.Fscan(in, &id1, &id2)

		id1, id2 = id1-1, id2-1

		if users[id1].Friends[id2] != nil || users[id2].Friends[id1] != nil {
			continue
		}

		users[id1].Friends[id2] = users[id2]
		users[id2].Friends[id1] = users[id1]

		users[id1].countOfFriends += 1
		users[id2].countOfFriends += 1
	}

	for i := range users {
		for j := 0; j < n; j++ {
			if users[i].Friends[j] == nil {
				continue
			}
			friendOfI := users[i].Friends[j]

			for k := range friendOfI.Friends {
				if k == j || k == i || friendOfI.Friends[k] == nil || areFriends(users[i], friendOfI.Friends[k]) {
					continue
				}

				mutualFriend(users[i], friendOfI.Friends[k])
			}
		}
	}

	for _, u := range users {
		if u.maxFriend == 0 {
			fmt.Fprint(out, 0, " ")
			fmt.Fprintln(out)
			continue
		}
		for uid, ff := range u.maxFriendWithUser {
			if ff == u.maxFriend {
				fmt.Fprint(out, uid+1, " ")
			}
		}
		fmt.Fprintln(out)
	}
}

func areFriends(u1, u2 *User) bool {
	return u1.Friends[u2.ID] != nil
}

func mutualFriend(u1, u2 *User) {
	u1.maxFriendWithUser[u2.ID] += 1
	u2.maxFriendWithUser[u1.ID] += 1

	if u1.maxFriend < u1.maxFriendWithUser[u2.ID] {
		u1.maxFriend = u1.maxFriendWithUser[u2.ID]
	}
	if u2.maxFriend < u2.maxFriendWithUser[u1.ID] {
		u2.maxFriend = u2.maxFriendWithUser[u1.ID]
	}
}
