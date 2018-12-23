package prison_after_ndays

func ConvertToInt(arr []int) int {
	ret := 0
	factor := 1
	for i := len(arr) - 1; i >= 0; i-- {
		ret += arr[i] * factor
		factor *= 2
	}

	return ret
}

func ConvertToArr(val int, length int) []int {
	ret := make([]int, length)
	remainder := val
	for i := length - 1; i >= 0; i-- {
		ret[i] = remainder % 2
		remainder /= 2
	}

	return ret
}

func PrisonAfterNDays(cells []int, N int) []int {
	curDay := 1
	curVal := ConvertToInt(cells)
	for curDay <= N {
		shiftRightVal := curVal >> 2

		curVal = (((curVal & 0x3F) ^ shiftRightVal) ^ 0x3F) << 1

		curDay++
	}

	ret := ConvertToArr(curVal, len(cells))
	return ret
}