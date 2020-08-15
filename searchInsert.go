package main

// func main() {
// 	fmt.Println(searchInsert([]int{1, 3}, 2))
// }

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1
	mid := 0
	for low < high {
		mid = low + ((high - low) / 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}
	if nums[mid] < target {
		for mid < len(nums) && nums[mid] < target {
			if mid+1 < len(nums) && nums[mid+1] > target {
				return mid + 1
			}
			mid++
		}
		return mid
	}
	for mid > 0 && nums[mid] > target {
		if mid-1 >= 0 && nums[mid-1] < target {
			return mid
		}
		mid--
	}
	return mid

}
