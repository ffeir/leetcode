package reverse_linked_list

import (
	"fmt"
	"testing"
)

func printLinkedList(head *Node) {
	curNode := head
	for curNode != nil {
		fmt.Printf("%d ", curNode.Val)
		curNode = curNode.Next
	}
}

func TestReverseLinkedList(t *testing.T) {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = &Node{Val: 4}

	// printLinkedList(head)

	newHead := ReverseLinkedList(head)
	printLinkedList(newHead)
}
