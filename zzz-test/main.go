package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode

	height int
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:    val,
			height: 1,
		}
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)
		root.height = Max(Height(root.Left), Height(root.Right)) + 1
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
		root.height = Max(Height(root.Left), Height(root.Right)) + 1
	} else {
		return root
	}

	balance := Height(root.Left) - Height(root.Right)

	if balance < -1 {
		if Height(root.Right.Right) < Height(root.Right.Left) {
			root.Right = LRotation(root.Right)
		}
		root = RRotation(root)
	}

	if balance > 1 {
		if Height(root.Left.Left) < Height(root.Left.Right) {
			root.Left = RRotation(root.Left)
		}
		root = LRotation(root)
	}

	updateHeights(root)

	return root
}

func LRotation(root *TreeNode) *TreeNode {
	var a, b, c *TreeNode
	a = root
	b = root.Left
	if b.Right != nil {
		c = b.Right
	}

	b.Right = a
	a.Left = c

	return b
}

func RRotation(root *TreeNode) *TreeNode {
	var a, b, c *TreeNode
	a = root
	b = root.Right
	if b.Left != nil {
		c = b.Left
	}

	b.Left = a
	a.Right = c

	return b
}

func updateHeights(root *TreeNode) int {
	if root == nil {
		return 0
	}
	root.height = Max(updateHeights(root.Left), updateHeights(root.Right)) + 1
	return root.height
}

func Print(root *TreeNode) {
	queue := make([]*TreeNode, 0, 8)
	queue = append(queue, root)

	level := 1

	for len(queue) > 0 {
		currLine := queue[:level]
		queue = queue[level:]
		level *= 2

		isNilLine := true
		for _, node := range currLine {
			if node == nil {
				fmt.Print(nil, " ")
				queue = append(queue, nil)
				queue = append(queue, nil)
			} else {
				if node.Left != nil || node.Right != nil {
					isNilLine = false
				}

				fmt.Print(node.Val, " ")
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}

		fmt.Println()
		if isNilLine {
			return
		}
	}
}

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.height
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func Delete(root *TreeNode, val int) {

// }

func main() {
	mass := []int{4, 1, 2, 3}

	root := (*TreeNode)(nil)
	for _, v := range mass {
		root = Insert(root, v)
	}

	Print(root)
}
