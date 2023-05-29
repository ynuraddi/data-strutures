package main

import (
	"container/heap"
	"fmt"
)

func main() {
	a := leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 3)
	fmt.Println(a)
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	hash := make([]int, 26)
	for _, v := range tasks {
		hash[v-'A']++
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, cnt := range hash {
		if cnt > 0 {
			heap.Push(h, cnt)
		}
	}

	time := 0
	queue := [][2]int{}
	for h.Len() > 0 || len(queue) > 0 {
		time++

		if h.Len() > 0 {
			taskCnt := heap.Pop(h).(int)
			if taskCnt > 1 {
				queue = append(queue, [2]int{taskCnt - 1, time + n})
			}
		}

		if len(queue) > 0 && queue[0][1] == time {
			taskCnt := queue[0][0]
			queue = queue[1:]
			if taskCnt > 0 {
				heap.Push(h, taskCnt)
			}
		}
	}

	return time
}
