package __reverse_integer

var (
	maxInt = 1 << 31 - 1
	minInt = -1 << 31
)

func Atoi(s string) int {
	ret := 0
	factor := 1

	endIndex := 0
	if s[0] == '-' {
		endIndex = 1
	}

	for i := len(s) - 1; i >= endIndex; i-- {
		ret += int(s[i] - '0') * factor
		factor *= 10
	}

	if endIndex == 1 {
		ret = -ret
	}
	return ret
}

func Itoa(x int) string {
	if x == 0 {
		return "0"
	}
	remainder := x
	runes := make([]rune, 12)
	i := 11

	neg := remainder < 0
	if neg {
		remainder = -remainder
	}

	for remainder > 0 {
		divisor := remainder % 10
		runes[i] = rune(divisor) + '0'
		i--
		remainder = remainder / 10
	}

	if neg {
		runes[i] = '-'
		i--
	}

	return string(runes[i + 1:12])
}

func Reverse(x int) int {
	var s string
	if x < 0 {
		s = Itoa(-x)
	} else {
		s = Itoa(x)
	}

	runes := []rune(s)
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	s = string(runes)
	if x < 0 {
		s = "-" + s
	}

	ret := Atoi(s)

	if (ret < 0 && ret < minInt) || (ret > 0 && ret > maxInt) {
		return 0
	}

	return ret
}
