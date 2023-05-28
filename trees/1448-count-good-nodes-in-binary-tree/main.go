package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	goodNodes(&TreeNode{})
}

func goodNodes(root *TreeNode) int {
	return goodNodesUtil(root, root.Val)
}

func goodNodesUtil(t *TreeNode, parent int) int {
	if t == nil {
		return 0
	}

	res := 1
	max := t.Val

	if parent > t.Val {
		res = 0
		max = parent
	}

	res += goodNodesUtil(t.Left, max)
	res += goodNodesUtil(t.Right, max)
	return res
}
