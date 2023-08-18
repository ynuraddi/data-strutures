package main

import (
	"container/heap"
	"math"
)

func main() {
	minimumDeviation([]int{4, 1, 5, 20, 3})
}

func minimumDeviation(nums []int) int {
	h := &IntHeap{}
	heap.Init(h)

	min := math.MaxInt

	for i := range nums {
		if nums[i]%2 == 0 {
			heap.Push(h, nums[i])
		} else {
			nums[i] *= 2
			heap.Push(h, nums[i])
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	ans := math.MaxInt
	for h.Top()%2 == 0 {
		max := heap.Pop(h).(int)

		if ans > max-min {
			ans = max - min
		}

		max /= 2

		if min > max {
			min = max
		}

		heap.Push(h, max)
	}

	if ans > h.Top()-min {
		ans = h.Top() - min
	}

	return ans
}

type IntHeap []int

func (h IntHeap) Top() int           { return h[0] }
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}
