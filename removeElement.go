package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] == val) {
		return 0
	}
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
