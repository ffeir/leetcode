package _22

func maxProfit(prices []int) int {
	max := 0

	if len(prices) == 0 || len(prices) == 1 {
		return max
	}

	bI := 0
	for i, v := range prices {
		if v >= prices[bI] {
			max += v - prices[bI]
		}

		bI = i
	}

	return max
}
