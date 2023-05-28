package main

func main() {
	levelOrder(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	r := make([][]int, 1)
	q := make([]*TreeNode, 1)
	level := 0

	q[0] = root

	for len(q) > 0 {
		currLen := len(q)
		for i := 0; i < currLen; i++ {
			node := q[0]
			q = q[1:]

			if len(r) <= level {
				r = append(r, []int{})
			}

			r[level] = append(r[level], node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}
	return r
}
