package main

import (
	"container/heap"
	"fmt"
)

func main() {
	t := swimInWater([][]int{
		{0, 1, 2, 3, 4},
		{24, 23, 22, 21, 5},
		{12, 13, 14, 15, 16},
		{11, 17, 18, 19, 20},
		{10, 9, 8, 7, 6},
	})
	fmt.Println(t)
}

func swimInWater(grid [][]int) (t int) {
	possibleWays := new(minHeap)
	heap.Init(possibleWays)

	visited := make(map[int]bool)
	n := len(grid)

	t = grid[0][0]
	heap.Push(possibleWays, position{0, 0, grid[0][0]})

	for t < n*n {
		currPosition := heap.Pop(possibleWays).(position)
		x, y := currPosition.x, currPosition.y
		visited[x*n+y] = true

		t = max(t, currPosition.height)
		if x == n-1 && y == n-1 {
			return t
		}

		if x+1 < n && !visited[(x+1)*n+y] {
			heap.Push(possibleWays, position{x + 1, y, grid[x+1][y]})
		}
		if x-1 >= 0 && !visited[(x-1)*n+y] {
			heap.Push(possibleWays, position{x - 1, y, grid[x-1][y]})
		}
		if y+1 < n && !visited[x*n+y+1] {
			heap.Push(possibleWays, position{x, y + 1, grid[x][y+1]})
		}
		if y-1 >= 0 && !visited[x*n+y-1] {
			heap.Push(possibleWays, position{x, y - 1, grid[x][y-1]})
		}
	}
	return
}

type position struct {
	x, y   int
	height int
}

type minHeap []position

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(position))
}

func (h *minHeap) Pop() interface{} {
	l := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return l
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
