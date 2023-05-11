package main

import (
	"fmt"
	"reflect"
)

func main() {
	test()
}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	prefix := 1
	for i := range nums {
		answer[i] = prefix
		prefix *= nums[i]
	}

	prefix = 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= prefix
		prefix *= nums[i]
	}

	return answer
}

func test() {
	cases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for _, c := range cases {
		if answer := productExceptSelf(c.nums); !reflect.DeepEqual(c.want, answer) {
			fmt.Println("error", answer, "!=", c.want)
			return
		}
	}
	fmt.Println("success")
}
