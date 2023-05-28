package main

func main() {
	isValidBST(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root, leftParent, rightParent *TreeNode) bool {
	if root == nil {
		return true
	}

	if leftParent != nil && root.Val <= leftParent.Val {
		return false
	}

	if rightParent != nil && root.Val >= rightParent.Val {
		return false
	}

	return isValid(root.Left, leftParent, root) && isValid(root.Right, root, rightParent)
}
