package main

import (
	"strings"
)

// func main() {
// 	fmt.Println(lengthOfLastWord("test test "))
// }

func lengthOfLastWord(s string) int {
	newStr := strings.TrimSpace(s)
	lenStr := len(newStr)
	if lenStr == 0 {
		return 0
	}
	var i int
	spaceExist := false
	for i = lenStr - 1; i > 0; i-- {
		if newStr[i] == 32 {
			spaceExist = true
			break
		}
	}
	if spaceExist {
		return lenStr - 1 - i
	}
	return lenStr
}
