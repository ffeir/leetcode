package __longest_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, LengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, LengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, LengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 2, LengthOfLongestSubstring("aab"))
	assert.Equal(t, 3, LengthOfLongestSubstring("dvdf"))
	assert.Equal(t, 5, LengthOfLongestSubstring("tmmzuxt"))
}
