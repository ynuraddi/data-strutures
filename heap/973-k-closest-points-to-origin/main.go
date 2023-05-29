package main

import "container/heap"

func main() {
}

type IntHeap [][]int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	return pow2(h[i][0])+pow2(h[i][1]) < pow2(h[j][0])+pow2(h[j][1])
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)   { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &IntHeap{}
	heap.Init(h)

	for i := range points {
		heap.Push(h, []int{points[i][0], points[i][1]})
	}

	var res [][]int
	for i := 0; i < k; i++ {
		tmp := heap.Pop(h).([]int)
		res = append(res, tmp)
	}

	return res
}

func pow2(a int) int {
	return a * a
}
