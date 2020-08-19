package main

import (
	"strings"
)

// func main() {
// 	fmt.Println(isPalindrome2("0P"))
// }
func isPalindrome2(s string) bool {
	lenStr := len(s)
	if lenStr == 0 {
		return true
	}
	i, j := 0, lenStr-1
	tempStr := strings.ToLower(s)
	for i <= j {
		if !((tempStr[i] >= 97 && tempStr[i] <= 122) || (tempStr[i] >= 48 && tempStr[i] <= 57)) {
			i++
			continue
		} else if !((tempStr[j] >= 97 && tempStr[j] <= 122) || (tempStr[j] >= 48 && tempStr[j] <= 57)) {
			j--
			continue
		}
		if tempStr[i] != tempStr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
