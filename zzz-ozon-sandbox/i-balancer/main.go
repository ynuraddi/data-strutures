package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const input = `
4 7
3 2 6 4
1 3
2 5
3 7
4 10
5 5
6 100
9 2

`

// Я РЕШАЮ ЗАДАЧИ В ШКОЛЕ, НО КОНТЕСТ ПОСТАРАЮСЬ ДОМА, У МЕНЯ НЕТ НОУТБУКА
func main() {
	// in := strings.NewReader(input)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	proc := &PowHeap{}
	heap.Init(proc)
	for i := 0; i < n; i++ {
		var pow int
		fmt.Fscan(in, &pow)
		heap.Push(proc, pow)
	}

	tasks := make([][2]int, m)
	for i := 0; i < m; i++ {
		var taskTime, taskDuration int
		fmt.Fscan(in, &taskTime, &taskDuration)
		tasks[i] = [2]int{taskTime, taskDuration}
	}

	var sum int
	time := &TimeHeap{}
	heap.Init(time)

	for _, t := range tasks {
		currentTime := t[0]
		for len(*time) > 0 {
			freeP := heap.Pop(time).([2]int)
			if freeP[1] > currentTime {
				heap.Push(time, freeP)
				break
			}
			heap.Push(proc, freeP[0])
		}

		if len(*proc) == 0 {
			continue
		}
		minP := heap.Pop(proc).(int)

		sum += t[1] * minP
		heap.Push(time, [2]int{minP, t[0] + t[1]})
	}

	fmt.Fprintln(out, sum)
}

type TimeHeap [][2]int

func (h TimeHeap) Len() int { return len(h) }
func (h TimeHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h TimeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TimeHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *TimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type PowHeap []int

func (h PowHeap) Len() int { return len(h) }
func (h PowHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h PowHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PowHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *PowHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
