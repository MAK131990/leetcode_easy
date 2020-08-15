package main

// func main() {
// 	fmt.Println(twoSum([]int{4, 2, 3}, 6))
// }

func twoSum(nums []int, target int) []int {

	mapArr := make(map[int]int)
	for index, value := range nums {
		val, ok := mapArr[value]
		if ok {
			return []int{val, index}
		}
		mapArr[target-value] = index
	}
	return nil
}
