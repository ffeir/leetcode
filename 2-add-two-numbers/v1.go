package __add_two_numbers

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func GetInteger(l *ListNode) int {
	res := 0
	curNode := l
	factor := 1
	for curNode != nil {
		res += curNode.Val * factor
		factor *= 10

		curNode = curNode.Next
	}

	return res
}

func Convert(input int) *ListNode {
	divisor := input

	var res *ListNode
	var lastNode *ListNode

	for ;; {
		remainder := divisor % 10
		divisor = divisor / 10
		if res == nil {
			res = &ListNode{remainder, nil}
			lastNode = res
		} else {
			lastNode.Next = &ListNode{remainder, nil}
			lastNode = lastNode.Next
		}

		if divisor == 0 {
			return res
		}
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return Convert(GetInteger(l1) + GetInteger(l2))
}
