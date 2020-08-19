package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{1, 2, 1, 4, 0}))
}

func checkPossibility(nums []int) bool {
	pos := -1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if pos != -1 {
				return false
			}
			pos = i - 1
		}
	}
	return pos == -1 || (pos == len(nums)-2) || pos == 0 || (nums[pos-1] <= nums[pos+1]) || nums[pos] <= nums[pos+2]

}
