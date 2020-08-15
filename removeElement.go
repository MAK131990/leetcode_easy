package main

// func main() {
// 	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
// }

func removeElement(nums []int, val int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] == val) {
		return 0
	}

	i := 0
	for i = 0; i < len(nums); i++ {
		if nums[i] != val {
			break
		}
	}
	if i == len(nums) {
		return 0
	}
	i = 0
	j := len(nums) - 1
	k := 0
	var temp int
	for i = 0; i <= j; i++ {
		if nums[i] == val {
			k++
			temp = nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j--
			i--
		}
	}
	if nums[j] == val {
		k++
	}
	return len(nums) - k
}
