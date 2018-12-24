package _61_repeated_element_in_2n_array

func RepeatedNTimes(A []int) int {
	l := len(A)
	N := l / 2
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		_, ok := m[A[i]]
		if ok {
			m[A[i]]++
			if m[A[i]] == N {
				return A[i]
			}
		} else {
			m[A[i]] = 1
		}
	}

	return -1
}
