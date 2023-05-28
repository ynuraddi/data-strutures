package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	rightSideView(&TreeNode{})
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var q []*TreeNode

	q = append(q, root)

	for len(q) > 0 {
		res = append(res, q[len(q)-1].Val)
		currLen := len(q)

		for i := 0; i < currLen; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return res
}
