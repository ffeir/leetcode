package main

// Time: O(n)
// Space: O(1)
func RemoveDuplicatedArrays(arrays []int64) int {
	if arrays == nil || len(arrays) == 0 {
		return 0
	}

	idx := 1
	for i := 1; i < len(arrays); i++ {
		if arrays[i] == arrays[idx-1] {
			continue
		}
		arrays[idx] = arrays[i]
		idx++
	}

	return idx
}
