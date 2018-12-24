package _61_repeated_element_in_2n_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatedNTimes(t *testing.T) {
	assert.Equal(t, 3, RepeatedNTimes([]int{1,2,3,3}))
	assert.Equal(t, 2, RepeatedNTimes([]int{2,1,2,5,3,2}))
	assert.Equal(t, 5, RepeatedNTimes([]int{5,1,5,2,5,3,5,4}))
}
