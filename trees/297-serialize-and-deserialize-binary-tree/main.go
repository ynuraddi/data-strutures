package main

import (
	"strconv"
	"strings"
)

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	res := make([]string, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "N")
			return
		}

		res = append(res, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	i := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if vals[i] == "N" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(vals[i])
		node := &TreeNode{Val: val}
		i++
		node.Left = dfs()
		node.Right = dfs()
		return node
	}

	return dfs()
}
