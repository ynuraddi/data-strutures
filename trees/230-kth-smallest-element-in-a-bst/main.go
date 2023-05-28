package main

func main() {
	kthSmallest(&TreeNode{}, 0)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0, k)

	for k != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}

	return -1
}