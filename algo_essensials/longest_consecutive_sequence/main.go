package longest_consecutive_sequence

func LongestConsecutiveSequence(arrays []int) int {
	if arrays == nil || len(arrays) == 0 {
		return 0
	}

	seqMap := make(map[int]struct{})
	for _, elem := range arrays {
		seqMap[elem] = struct{}{}
	}

	for _, elem := range arrays {
		i := elem
		for {
			i--
			_, ok := seqMap[i]
			if !ok {
				break
			}
		}
	}
}
