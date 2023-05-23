package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	listsLen := len(lists)
	if listsLen == 0 {
		return nil
	} else if listsLen == 1 {
		return lists[0]
	}

	for len(lists) > 1 {
		list1, list2 := lists[0], lists[1]
		merged := mergeLists(list1, list2)

		lists = lists[2:]
		lists = append(lists, merged)
	}
	return lists[0]
}

func mergeLists(listA, listB *ListNode) *ListNode {
	res := &ListNode{}
	dummyHead := res
	for listA != nil || listB != nil {
		if listA == nil {
			res.Next = listB
			break
		}
		if listB == nil {
			res.Next = listA
			break
		}

		if listA.Val < listB.Val {
			res.Next = listA
			res = res.Next
			listA = listA.Next
		} else {
			res.Next = listB
			res = res.Next
			listB = listB.Next
		}
	}
	return dummyHead.Next
}
