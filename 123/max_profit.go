package _23

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	max1, max2 := 0, 0
	tempMax := 0

	bI := 0 // buy index
	for i, v := range prices {
		if v >= prices[bI] {
			tempMax += v - prices[bI]
		} else { // new buy txn
			if max2 < tempMax {
				max1 = max2
				max2 = tempMax
			}
			tempMax = 0
			fmt.Println(max1, max2, tempMax)
		}

		bI = i
	}

	if tempMax > max1 {
		max1 = tempMax
	}

	return max1 + max2
}
