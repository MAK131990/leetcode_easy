package main

// func main() {
// 	fmt.Println(romanToInt("MCMXCIV"))
// }

func romanToInt(s string) int {

	charMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for i := 0; i < len(s); {
		if i+1 == len(s) {
			result += charMap[string(s[i])]
			i++
		} else {
			cur := string(s[i])
			next := string(s[i+1])
			if (cur == "I" && (next == "V" || next == "X")) ||
				(cur == "X" && (next == "L" || next == "C")) ||
				(cur == "C" && (next == "D" || next == "M")) {
				result = result + charMap[next] - charMap[cur]
				i = i + 2
			} else {
				result = result + charMap[cur]
				i++
			}
		}
	}
	return result
}
