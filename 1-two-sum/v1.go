package two_sum

func twoSumV1(nums []int, target int) []int {
	for i, v := range nums {
		for j := len(nums) - 1; j > i; j-- {
			sum := nums[j] + v
			if sum == target {
				return []int {i, j}
			}
		}
	}

	return []int {-1, -1}
}