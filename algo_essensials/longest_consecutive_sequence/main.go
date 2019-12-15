package longest_consecutive_sequence

func LongestConsecutiveSequence(arrays []int) int {
	if arrays == nil || len(arrays) == 0 {
		return 0
	}

	seqMap := make(map[int]struct{})
	for _, elem := range arrays {
		seqMap[elem] = struct{}{}
	}

	maxLength := 1
	for _, elem := range arrays {
		i := elem
		length := 1
		for {
			i--
			_, ok := seqMap[i]
			if !ok {
				break
			}
			length++
		}

		j := elem
		for {
			j++
			_, ok := seqMap[j]
			if !ok {
				break
			}
			length++
		}

		if maxLength < length {
			maxLength = length
		}
	}

	return maxLength
}
