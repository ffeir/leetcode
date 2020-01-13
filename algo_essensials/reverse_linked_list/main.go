package reverse_linked_list

type Node struct {
	Val  int
	Next *Node
}

func ReverseLinkedList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var tail *Node
	newHead := head.Next
	curNode := head

	for newHead != nil {
		oldNext := newHead.Next

		curNode.Next = tail
		newHead.Next = curNode

		tail = curNode
		curNode = newHead
		newHead = oldNext
	}

	return newHead
}
