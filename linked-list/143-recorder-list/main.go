package main

import "fmt"

func main() {
	var (
		l5 = &ListNode{5, nil}
		l4 = &ListNode{4, l5}
		l3 = &ListNode{3, l4}
		l2 = &ListNode{2, l3}
		l1 = &ListNode{1, l2}
	)

	reorderList(l1)

	for l1 != nil {
		fmt.Println(l1)
		l1 = l1.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	list := make([]*ListNode, 0, 10)

	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	i, j := 0, len(list)-1

	for ; i < j; i, j = i+1, j-1 {
		list[j].Next = list[i].Next
		list[i].Next = list[j]
	}

	list[i].Next = nil
}
