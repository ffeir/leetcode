package two_sum

func twoSumV2(nums []int, target int) []int {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}

	for i, v := range nums {
		anotherValue := target - v
		if anotherValue == v {
			if len(m[anotherValue]) > 1 {
				return m[v]
			}
			continue
		}

		anotherIndex, ok := m[anotherValue]
		if ok {
			return []int {i, anotherIndex[0]}
		}
	}

	return []int {-1, -1}
}


