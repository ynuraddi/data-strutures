package main

func main() {
	var (
		l5 = &ListNode{5, nil}
		l4 = &ListNode{4, l5}
		l3 = &ListNode{3, l4}
		l2 = &ListNode{2, l3}
		l1 = &ListNode{1, l2}
	)

	reverseList(l1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head, prev, head.Next = head.Next, head, prev
	}
	return prev
}
