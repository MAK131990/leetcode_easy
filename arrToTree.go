package main

import "fmt"

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

func sortedArrayToBST(nums []int) *TreeNode {
	var res TreeNode
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		res.Val = nums[0]
		res.Left, res.Right = nil, nil
		return &res
	}
	mid := len(nums) / 2
	res.Val = nums[mid]
	res.Left = sortedArrayToBST(nums[0:mid])
	res.Right = sortedArrayToBST(nums[mid+1:])

	return &res
}
