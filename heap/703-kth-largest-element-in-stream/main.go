package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(h.Add(3))
	fmt.Println(h.Add(5))
	fmt.Println(h.Add(10))
	fmt.Println(h.Add(9))
	fmt.Println(h.Add(4))
	fmt.Println(*h.minHeap)
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	tmp := IntHeap(nums)
	this := KthLargest{minHeap: &tmp, k: k}
	heap.Init(this.minHeap)
	for len(*this.minHeap) > k {
		heap.Pop(this.minHeap)
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	if len(*this.minHeap) > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}

// type IntHeap []int

// func (h IntHeap) Get(i int) int      { return h[i] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
// func (h IntHeap) Max(i, j int) int {
// 	if h[i] > h[j] {
// 		return i
// 	}
// 	return j
// }

// func (h *IntHeap) Push(x int) {
// 	*h = append(*h, x)
// 	xindex := len(*h) - 1
// 	for xindex >= 0 && h.Less(xindex, (xindex-1)/2) {
// 		h.Swap(xindex, (xindex-1)/2)
// 		xindex = (xindex - 1) / 2
// 	}
// }

// func (h *IntHeap) Pop() {
// 	(*h)[0] = (*h)[len(*h)-1]
// 	*h = (*h)[0 : len(*h)-1]
// 	xindex := 0
// 	for xindex < len(*h) && xindex*2+1 < len(*h) && xindex*2+2 < len(*h) {
// 		max := h.Max(xindex*2+1, xindex*2+2)
// 		if h.Less(max, xindex) {
// 			h.Swap(max, xindex)
// 		}
// 		xindex = max
// 	}
// }

// type KthLargest struct {
// 	minHeap *IntHeap
// 	k       int
// }

// func Constructor(k int, nums []int) KthLargest {
// 	heap := KthLargest{minHeap: &IntHeap{}, k: k}
// 	for i := range nums {
// 		heap.minHeap.Push(nums[i])
// 	}
// 	return heap
// }

// func (this *KthLargest) Add(val int) int {
// 	this.minHeap.Push(val)
// 	k := this.k
// 	h := &IntHeap{}
// 	for i := range *this.minHeap {
// 		h.Push((*this.minHeap)[i])
// 	}
// 	for len(*h) > 0 && k > 1 {
// 		h.Pop()
// 		k--
// 	}
// 	return (*h)[0]
// }
