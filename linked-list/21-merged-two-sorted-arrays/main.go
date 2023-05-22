package main

func main() {
	var (
		l3 = &ListNode{3, nil}
		l2 = &ListNode{2, l3}
		l1 = &ListNode{1, l2}

		l32 = &ListNode{5, nil}
		l22 = &ListNode{4, l32}
		l12 = &ListNode{2, l22}
	)

	mergeTwoLists(l1, l12)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var head *ListNode
	var t *ListNode

	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}

	t = list1
	list1 = list1.Next
	head = t

	for t != nil {
		if list1 == nil && list2 != nil || list2 != nil && list1.Val >= list2.Val {
			t.Next = list2
			t, list2 = t.Next, list2.Next
		} else if list2 == nil && list1 != nil || list1 != nil && list1.Val <= list2.Val {
			t.Next = list1
			t, list1 = t.Next, list1.Next
		}
	}

	return head
}
