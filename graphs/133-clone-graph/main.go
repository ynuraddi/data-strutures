package main

import "fmt"

func main() {
	var (
		a = &Node{Val: 1, Neighbors: []*Node{}}
		b = &Node{Val: 2, Neighbors: []*Node{}}
		c = &Node{Val: 3, Neighbors: []*Node{}}
		d = &Node{Val: 4, Neighbors: []*Node{}}
	)
	a.Neighbors = append(a.Neighbors, b, d)
	b.Neighbors = append(b.Neighbors, a, c)
	c.Neighbors = append(c.Neighbors, b, d)
	d.Neighbors = append(d.Neighbors, a, c)

	aboba := cloneGraph(a)
	fmt.Println(aboba)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[*Node]*Node{}

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if _, ok := visited[node]; ok {
			return visited[node]
		}

		newNode := &Node{Val: node.Val}
		visited[node] = newNode

		for _, n := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(n))
		}
		return newNode
	}
	return dfs(node)
}
