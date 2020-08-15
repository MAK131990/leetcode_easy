package main

// func main() {
// 	fmt.Println(plusOne([]int{0}))
// }

func plusOne(digits []int) []int {
	lenDig := len(digits)
	for i := lenDig - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 {
			digits[i] %= 10
		} else {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
