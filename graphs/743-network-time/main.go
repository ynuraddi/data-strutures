package main

import "container/heap"

func main() {
	// networkDelayTime([][]int{
	// 	{1, 2, 1},
	// 	{2, 1, 3},
	// }, 2, 2)
	networkDelayTime([][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
		{2, 4, 5},
	}, 4, 2)
}

func networkDelayTime(times [][]int, n int, k int) (answer int) {
	graph := make(map[int][]neighbour)
	for _, log := range times {
		graph[log[0]] = append(graph[log[0]], neighbour{log[1], log[2]})
	}

	h := &minHeap{heapNode{distance: 0, nodeIndex: k}}
	heap.Init(h)

	visited := make(map[int]bool)
	t := 0

	for !h.isEmpty() {
		hNode := heap.Pop(h).(heapNode)

		if vis := visited[hNode.nodeIndex]; vis {
			continue
		}
		visited[hNode.nodeIndex] = true

		t = max(t, hNode.distance)

		neighbours := graph[hNode.nodeIndex]
		for _, neigh := range neighbours {
			if vis := visited[neigh.destination]; !vis {
				heap.Push(h, heapNode{
					distance:  neigh.weight + hNode.distance,
					nodeIndex: neigh.destination,
				})
			}
		}
	}

	if n == len(visited) {
		return t
	}
	return -1
}

type neighbour struct {
	destination int
	weight      int
}

type heapNode struct {
	distance  int
	nodeIndex int
}

type minHeap []heapNode

func (h minHeap) Len() int           { return len(h) }
func (h *minHeap) isEmpty() bool     { return len(*h) == 0 }
func (h minHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(heapNode))
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
