package main

import (
	"strconv"
)

// func main() {
// 	fmt.Println(addBinary("1111", "100"))
// }

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}
	var longerStringLen int
	if len(a) > len(b) {
		longerStringLen = len(a)
	} else {
		longerStringLen = len(b)
	}
	result := make([]int, longerStringLen+1)
	var carry = 0
	k := longerStringLen
	var i, j int
	for i, j = len(a)-1, len(b)-1; i >= 0 && j >= 0; {
		first, _ := strconv.Atoi(string(a[i]))
		second, _ := strconv.Atoi(string(b[j]))
		sum := first + second + carry
		result[k] = (sum % 2)
		carry = sum / 2
		i--
		j--
		k--
	}
	if carry == 1 && i < 0 && j < 0 {
		result[k] = 1
	} else {
		if i >= 0 {
			for i >= 0 {
				first, _ := strconv.Atoi(string(a[i]))
				sum := first + carry
				result[k] = (sum % 2)
				carry = sum / 2
				i--
				k--
			}
		} else if j >= 0 {
			for j >= 0 {
				first, _ := strconv.Atoi(string(b[j]))
				sum := first + carry
				result[k] = (sum % 2)
				carry = sum / 2
				j--
				k--
			}
		}
	}
	if carry == 1 {
		result[k] = 1
	}
	resultStr := ""
	for i = 0; i < len(result); i++ {
		if resultStr == "" && result[i] == 0 {
			continue
		}
		resultStr += strconv.Itoa(result[i])
	}
	if len(resultStr) == 0 && len(result) > 0 {
		return "0"
	}
	return resultStr
}
