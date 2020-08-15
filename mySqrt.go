package main

import "fmt"

func main() {
	fmt.Println(mySqrt(2))
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	var mid int = x / 2
	var sq int
	var sq1 int
	for i := 1; i <= mid; i++ {
		sq = i * i
		sq1 = (i + 1) * (i + 1)
		if sq == x {
			return i
		} else if sq < x && sq1 > x {
			return i
		} else if sq < x && sq1 == x {
			return i + 1
		}
	}
	return 0
}
