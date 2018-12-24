package _62_maximum_width_ramp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxWidthRamp(t *testing.T) {
	assert.Equal(t, 4, MaxWidthRamp([]int{6,0,8,2,1,5}))
	assert.Equal(t, 7, MaxWidthRamp([]int{9,8,1,0,1,9,4,0,4,1}))
}
