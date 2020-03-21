package main

func RemoveDuplicateArray(arrays []int) int {
	if arrays == nil || len(arrays) == 0 {
		return 0
	}

	idx := 1
	dupCount := 1
	for i := 1; i < len(arrays); i++ {
		if arrays[i] == arrays[idx-1] {
			if dupCount >= arrays[idx-1] {
				continue
			}
			dupCount++
		} else {
			dupCount = 1
		}

		arrays[idx] = arrays[i]
		idx++
	}

	return idx
}
