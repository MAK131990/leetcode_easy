package main

import (
	"math"
)

// func main() {
// 	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
// }

func maxSubArray(nums []int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	} else if lenNums == 1 {
		return nums[0]
	}
	maxSoFar := math.MinInt32
	maxEndingHere := 0
	for i := 0; i < lenNums; i++ {
		maxEndingHere += nums[i]
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}
