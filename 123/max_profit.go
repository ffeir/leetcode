package _23

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	lp1 := 1<<63 - 1 // lowest price in transaction 1
	m1 := 0          // max profit in transaction 1
	lp2 := lp1
	m2 := 0 // max profit in transaction 2

	for _, v := range prices {
		lp1 = min(lp1, v)
		m1 = max(m1, v-lp1)
		lp2 = min(v)
	}
}

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
