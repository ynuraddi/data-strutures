package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	stack := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := stack[head]; ok {
			return true
		}
		stack[head] = struct{}{}
		head = head.Next
	}

	return false
}
