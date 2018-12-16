package __reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, 123, Reverse(321))
	assert.Equal(t, -123, Reverse(-321))
	assert.Equal(t, 0, Reverse(math.MaxInt32))
	assert.Equal(t, 0, Reverse(math.MinInt32))
	assert.Equal(t, 0, Reverse(2147483648))
	assert.Equal(t, 0, Reverse(-9147483648))
	assert.Equal(t, 0, Reverse(0))

	assert.Equal(t, 0, Reverse(1534236469))
}

func TestAtoi(t *testing.T) {
	assert.Equal(t, 123, Atoi("123"))
	assert.Equal(t, -123, Atoi("-123"))
	assert.Equal(t, math.MinInt32, Atoi(strconv.Itoa(math.MinInt32)))
	assert.Equal(t, math.MaxInt32, Atoi(strconv.Itoa(math.MaxInt32)))
}

func TestItoa(t *testing.T) {
	assert.Equal(t, "123", Itoa(123))
	assert.Equal(t, "-123", Itoa(-123))

	assert.Equal(t, "2147483647", Itoa(math.MaxInt32))
	assert.Equal(t, "-2147483648", Itoa(math.MinInt32))
}
