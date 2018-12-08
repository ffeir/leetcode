package __add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestAddTwoNumbers(t *testing.T) {
	l := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	assert.Equal(t, 807, GetInteger(AddTwoNumbers(l, l2)))

	l = &ListNode{5, nil}
	l2 = &ListNode{5, nil}
	assert.Equal(t, 10, GetInteger(AddTwoNumbers(l, l2)))
}
