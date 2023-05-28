package main

func main() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	mid := index(inorder, preorder[0])
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}

func index(arr []int, k int) int {
	for i, v := range arr {
		if v == k {
			return i
		}
	}
	return -1
}
