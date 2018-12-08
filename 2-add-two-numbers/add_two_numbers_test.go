package __add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInteger(t *testing.T) {
	l := &ListNode{0, nil}
	assert.Equal(t, 0, GetInteger(l))

	l.Next = &ListNode{2, nil}
	assert.Equal(t, 20, GetInteger(l))

	l.Next.Next = &ListNode{8, nil}
	assert.Equal(t, 820, GetInteger(l))
}

func TestConvert(t *testing.T) {
	l := &ListNode{0, nil}
	l.Next = &ListNode{2, nil}
	l.Next.Next = &ListNode{8, nil}

	assert.Equal(t, 820, GetInteger(Convert(820)))
	assert.Equal(t, 342, GetInteger(Convert(342)))
	assert.Equal(t, 465, GetInteger(Convert(465)))
}

func TestAddTwoNumbers(t *testing.T) {
	l := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	assert.Equal(t, 807, GetInteger(AddTwoNumbers(l, l2)))
}
