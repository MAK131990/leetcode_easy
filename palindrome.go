package main

import "fmt"

// func main() {
// 	fmt.Println(isPalindrome(101))
// }

func isPalindrome(x int) bool {
	// tempVar := x
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	var result int = 0
	// var dig int
	for result < x {
		// dig = tempVar % 10
		result = (result * 10) + (x % 10)
		x = x / 10
		fmt.Println(x, result)
	}
	fmt.Println(x, result)
	return result == x || result/10 == x
}
