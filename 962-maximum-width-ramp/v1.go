package _62_maximum_width_ramp

func MaxWidthRamp(A []int) int {
	maxWidth := 0
	for i := 0; i < len(A) - 1; i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] <= A[j] {
				if maxWidth < j - i {
					maxWidth = j - i
				}
			}
		}
	}

	return maxWidth
}
