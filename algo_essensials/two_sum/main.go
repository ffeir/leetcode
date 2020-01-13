package two_sum

func TwoSum(arrays []int, target int) (int, int) {
	m := make(map[int]int)

	for i, elem := range arrays {
		if elem > target {
			continue
		}

		m[elem] = i
	}

	for idx1, k := range m {
		idx2, ok := m[target-k]
		if ok {
			if idx1 < idx2 {
				return idx1, idx2
			} else {
				return idx2, idx1
			}
		}
	}

	return 0, 0
}
