package main

import (
	"fmt"
	"reflect"
)

func main() {
	test()
}

func longestConsecutive(nums []int) int {
	h := make(map[int]int)
	longest := 0

	for _, n := range nums {
		_, ok := h[n]
		if !ok {
			left := h[n-1]
			right := h[n+1]
			h[n] = 1 + left + right
			if left > 0 {
				h[n-left] = h[n]
			}
			if right > 0 {
				h[n+right] = h[n]
			}
			if longest < h[n] {
				longest = h[n]
			}
		}
	}

	return longest
}

func test() {
	cases := []struct {
		nums []int
		want int
	}{
		// {
		// 	nums: []int{0, -1},
		// 	want: 2,
		// },
		{
			nums: []int{100, 4, 200, 1, 3, 2, 9, 6, 7, 8, 5},
			want: 9,
		},
	}

	for _, c := range cases {
		if answer := longestConsecutive(c.nums); !reflect.DeepEqual(c.want, answer) {
			fmt.Println("error", answer, "!=", c.want)
			return
		}
	}
	fmt.Println("success")
}
