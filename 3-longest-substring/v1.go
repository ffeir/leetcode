package __longest_substring

func LengthOfLongestSubstring(s string) int {
	maxLen := 0

	subElem := make(map[uint8]int)

	startIndex := 0

	dupInLast := false

	for i := 0; i < len(s); i++ {
		e := s[i]

		eIndex, ok := subElem[e]

		if eIndex < startIndex {
			ok = false
		}

		if ok {
			curLen := i - startIndex
			if maxLen < curLen {
				maxLen = curLen
			}

			startIndex = subElem[e] + 1

			dupInLast = i == len(s) - 1
		}

		subElem[e] = i
	}

	if !dupInLast {
		lastLen := len(s) - startIndex
		if maxLen < lastLen {
			maxLen = lastLen
		}
	}

	return maxLen
}