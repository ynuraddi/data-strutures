package main

import (
	"fmt"
	"math"
)

func main() {
	a := maxPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -2,
			},
		},
	})
	maxPathSum(&TreeNode{Val: -3})
	maxPathSum(&TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	})

	fmt.Println(a)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	globalMax := math.MinInt32
	dfs(root, &globalMax)
	return globalMax
}

func dfs(root *TreeNode, globalMax *int) int {
	if root == nil {
		return 0
	}

	pathSumFromLeft := max(dfs(root.Left, globalMax), 0)
	pathSumFromRight := max(dfs(root.Right, globalMax), 0)

	*globalMax = max(*globalMax, root.Val+pathSumFromLeft+pathSumFromRight)

	return root.Val + max(pathSumFromLeft, pathSumFromRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
