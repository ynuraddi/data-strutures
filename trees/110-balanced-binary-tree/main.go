package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return dfs(root) >= 0
}

func dfs(t *TreeNode) int {
	if t == nil {
		return 0
	}

	left := dfs(t.Left)
	right := dfs(t.Right)

	if right == left || right+1 == left || left+1 == right {
		return max(left, right) + 1
	}
	return -10
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
