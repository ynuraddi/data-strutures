package main

func main() {
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	curr := head
	for curr != nil {
		newNode := &Node{
			Val:    curr.Val,
			Next:   curr.Next,
			Random: curr.Random,
		}

		curr.Next = newNode
		curr = newNode.Next
	}

	curr = head
	for curr != nil {
		curr = curr.Next
		if curr.Random != nil {
			curr.Random = curr.Random.Next
		}
		curr = curr.Next
	}

	dummy := &Node{}
	curr2 := dummy
	curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = curr.Next.Next
		curr = curr.Next

		curr2.Next = next
		curr2 = curr2.Next
	}

	return dummy.Next
}
