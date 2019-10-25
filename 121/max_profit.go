package _21

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	lowestPrice := 1<<63 - 1
	m := -1 << 63 // max profit
	for _, v := range prices {
		lowestPrice = min(lowestPrice, v)
		m = max(m, v-lowestPrice)
	}

	return m
}
