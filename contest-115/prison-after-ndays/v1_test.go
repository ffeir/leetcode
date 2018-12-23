package prison_after_ndays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToInt(t *testing.T) {
	assert.Equal(t, 1, ConvertToInt([]int{0, 0, 1}))
	assert.Equal(t, 3, ConvertToInt([]int{0, 1, 1}))
	assert.Equal(t, 7, ConvertToInt([]int{1, 1, 1}))
	assert.Equal(t, 6, ConvertToInt([]int{1, 1, 0}))
	assert.Equal(t, 96, ConvertToInt([]int{0, 1, 1, 0, 0, 0, 0, 0}))
}

func TestConvertToArr(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 0, 0, 0, 0, 0}, ConvertToArr(ConvertToInt([]int{0, 1, 1, 0, 0, 0, 0, 0}), 8))
}

func TestPrisonAfterNDays(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 0, 0, 0, 0, 0}, PrisonAfterNDays([]int{0,1,0,1,1,0,0,1}, 1))
	assert.Equal(t, []int{0,0,1,1,0,0,0,0}, PrisonAfterNDays([]int{0,1,0,1,1,0,0,1}, 7))
	assert.Equal(t, []int{0,0,1,1,1,1,1,0}, PrisonAfterNDays([]int{1,0,0,1,0,0,1,0}, 1000000000))
	PrisonAfterNDays([]int{1,1,1,1,1,1,0,0}, 911783319)
}
