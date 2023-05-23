package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	arr := []*ListNode{}
	h := head

	i := 0
	for h != nil && i < k {
		arr = append(arr, h)
		h = h.Next
		i++
	}

	if h == nil && i != k {
		return head
	}

	for i := range arr {
		if i == 0 {
			arr[i].Next = reverseKGroup(arr[len(arr)-1].Next, k)
		} else {
			arr[i].Next = arr[i-1]
		}
	}

	return arr[len(arr)-1]
}
