package main

func main() {
	var (
		l3 = &ListNode{Val: 2, Next: nil}
		l2 = &ListNode{Val: 3, Next: l3}
		l1 = &ListNode{Val: 8, Next: l2}

		b3 = &ListNode{Val: 1, Next: nil}
		b2 = &ListNode{Val: 2, Next: b3}
		b1 = &ListNode{Val: 9, Next: b2}
	)

	addTwoNumbers(l1, b1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	d := res

	val := 0
	for {
		nL := new(ListNode)
		sum := l1.Val + l2.Val + val
		val = sum / 10
		nL.Val = sum % 10

		res.Next = nL
		res = res.Next

		if l1.Next == nil && l2.Next == nil && val == 0 {
			return d.Next
		}

		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}
}
