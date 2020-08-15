package main

import (
	"math"
)

// func main() {
// 	fmt.Println(reverse(-123))
// }

func reverse(x int) int {
	isNegative := false
	tempVar := x

	if x < 0 {
		isNegative = true
		tempVar = -x
	}

	newNum := 0
	for tempVar > 0 {
		digit := tempVar % 10
		if !isNegative && (newNum > math.MaxInt32/10 || newNum*10 > math.MaxInt32-digit) {
			return 0
		} else if isNegative && (-newNum < math.MinInt32/10 || -newNum*10 < math.MinInt32+digit) {
			return 0
		}
		newNum = newNum*10 + digit
		tempVar /= 10
	}
	if isNegative {
		newNum = -newNum
	}
	return newNum
}
