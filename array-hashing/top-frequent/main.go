package main

import (
	"fmt"
	"reflect"
)

func main() {
	test()
}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int, len(nums))
	max := 0

	for _, v := range nums {
		hashMap[v]++
		if max < hashMap[v] {
			max = hashMap[v]
		}
	}

	hash := make([][]int, max+1)

	for value, count := range hashMap {
		hash[count] = append(hash[count], value)
	}

	res := make([]int, 0, k)
	c := k
	for i := len(hash) - 1; i > 0 && c > 0; i-- {
		if len(hash[i]) != 0 {
			if len(hash[i]) < c {
				res = append(res, hash[i]...)
			} else {
				res = append(res, hash[i][:c]...)
			}
			c -= len(hash[i])
		}
	}

	return res
}

func test() {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 6},
			k:    3,
			want: []int{4, 1},
		},
	}

	for _, c := range cases {
		if answer := topKFrequent(c.nums, c.k); !reflect.DeepEqual(c.want, answer) {
			fmt.Println("error", answer, "!=", c.want)
			return
		}
	}
	fmt.Println("success")
}
