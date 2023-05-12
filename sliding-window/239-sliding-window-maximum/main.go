package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 0 {
		return []int{}
	} else if k == 1 {
		return nums
	}

	var (
		res = make([]int, len(nums)-k+1)
		dq  = new(deque)
	)

	for i := range nums {
		if !dq.empty() && dq.getFirst() == i-k {
			dq.popFirst()
		}

		for !dq.empty() && nums[dq.getLast()] < nums[i] {
			dq.popLast()
		}

		dq.push(i)

		if i >= k-1 {
			res[i-k+1] = nums[dq.getFirst()]
		}
	}

	return res
}

type deque struct {
	indexes []int
}

func (d *deque) push(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *deque) getFirst() int {
	return d.indexes[0]
}

func (d *deque) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *deque) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *deque) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *deque) empty() bool {
	return len(d.indexes) == 0
}
