package main

// func main() {
// 	fmt.Println(strStr("mississippi", "pi"))
// }

func strStr(haystack string, needle string) int {
	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle == 0 {
		return 0
	}
	if lenNeedle > lenHaystack {
		return -1
	}
	if lenNeedle == lenHaystack && needle != haystack {
		return -1
	} else if lenNeedle == lenHaystack && needle == haystack {
		return 0
	}
	for i := 0; i+lenNeedle <= lenHaystack; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}
	return -1
}
