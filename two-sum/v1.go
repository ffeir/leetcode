package main

import "fmt"

func twoSum(nums []int, target int) []int {
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

func main()  {
	fmt.Println(twoSum([]int {3,2,4}, 6))
	fmt.Println(twoSum([]int {3,3}, 6))
}
