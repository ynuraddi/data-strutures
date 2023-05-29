package main

import (
	"container/heap"
)

func main() {
	lastStoneWeight([]int{9, 10, 4, 5, 7, 1})
}

type IntHeap []int

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

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := range stones {
		heap.Push(h, stones[i])
	}

	for len(*h) > 1 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		if a != b {
			heap.Push(h, a-b)
		}
	}

	if len(*h) == 0 {
		return 0
	}
	return h.Pop().(int)
}
