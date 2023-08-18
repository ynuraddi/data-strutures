package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	r := root

	prev := root
	for root != nil {
		prev = root

		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	if val < prev.Val {
		prev.Left = &TreeNode{
			Val: val,
		}
	} else {
		prev.Right = &TreeNode{
			Val: val,
		}
	}

	return r
}
