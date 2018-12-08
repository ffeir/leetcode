package __add_two_numbers

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var lastNode *ListNode

	lastNodeInL1 := l1
	lastNodeInL2 := l2

	gt10 := false
	for lastNodeInL1 != nil || lastNodeInL2 != nil {
		s := 0
		if lastNodeInL1 != nil {
			s += lastNodeInL1.Val

			lastNodeInL1 = lastNodeInL1.Next
		}

		if lastNodeInL2 != nil {
			s += lastNodeInL2.Val

			lastNodeInL2 = lastNodeInL2.Next
		}

		if gt10 {
			s += 1
		}

		curNodeValue := 0
		if s >= 10 {
			curNodeValue = s - 10
			gt10 = true
		} else {
			curNodeValue = s
			gt10 = false
		}

		if res == nil {
			res = &ListNode{curNodeValue, nil}
			lastNode = res
		} else {
			lastNode.Next = &ListNode{curNodeValue, nil}
			lastNode = lastNode.Next
		}
	}

	// l1 and l2 have the same length and the sum of last Node larger than 10
	if gt10 {
		lastNode.Next = &ListNode{1, nil}
	}

	return res
}