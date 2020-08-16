package main

// func main() {
// 	fmt.Println(climbStairs(4))
// }

var memo map[int]int = make(map[int]int)

func climbStairs(n int) int {
	res, ok := memo[n]
	if ok {
		return res
	}
	if n <= 2 {
		memo[n] = n
		return n
	}
	memo[n] = climbStairs(n-1) + climbStairs(n-2)
	return memo[n]
}
