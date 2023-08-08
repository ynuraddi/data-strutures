package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type developer struct {
	index int
	level int
	abs   int
}

// const input = `
// 3
// 6
// 2 1 3 1 1 4
// 2
// 5 5
// 8
// 1 4 2 5 4 2 6 3
// `

func main() {
	in := bufio.NewReader(os.Stdin)
	// in := strings.NewReader(input)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var l int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l)

		arr := make([]developer, l)
		arrHeaps := []*DevHeap{}

		for j := 0; j < l; j++ {
			var d developer
			d.index = j + 1
			fmt.Fscan(in, &d.level)

			arr[j] = d

			for q := range arrHeaps {
				dCopy := d
				dCopy.abs = abs(arr[j].level, arr[q].level)
				heap.Push(arrHeaps[q], &dCopy)
			}
			arrHeaps = append(arrHeaps, &DevHeap{})
		}

		usedDevs := make(map[int]bool, l)

		var first, second developer
		for j := 0; j < l; j++ {
			if usedDevs[arr[j].index] {
				continue
			}
			first = arr[j]
			usedDevs[first.index] = true

			for arrHeaps[j].Len() > 0 {
				second = *heap.Pop(arrHeaps[j]).(*developer)
				if !usedDevs[second.index] {
					break
				}
			}
			usedDevs[second.index] = true

			fmt.Fprintln(out, first.index, second.index)
		}
		fmt.Fprintln(out)
	}
}

func abs(a, b int) int {
	res := a - b

	if res < 0 {
		return -res
	}
	return res
}

type DevHeap []*developer

func (h DevHeap) Len() int { return len(h) }
func (h DevHeap) Less(i, j int) bool {
	if h[i].abs == h[j].abs {
		return h[i].index < h[j].index
	}
	return h[i].abs < h[j].abs
}
func (h DevHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *DevHeap) Push(x interface{}) { *h = append(*h, x.(*developer)) }
func (h *DevHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
