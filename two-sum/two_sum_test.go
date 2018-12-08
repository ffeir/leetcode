package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumV1(t *testing.T)  {
	assert.Equal(t, []int {1, 2}, twoSumV1([]int {3,2,4}, 6))
	assert.Equal(t, []int {0, 1}, twoSumV1([]int {3,3}, 6))
	assert.Equal(t, []int {0, 1}, twoSumV1([]int {2, 7, 11, 15}, 9))
}

func TestTwoSumV2(t *testing.T) {
	assert.Equal(t, []int {1, 2}, twoSumV2([]int {3,2,4}, 6))
	assert.Equal(t, []int {0, 1}, twoSumV2([]int {3,3}, 6))
	assert.Equal(t, []int {0, 1}, twoSumV2([]int {2, 7, 11, 15}, 9))
}
