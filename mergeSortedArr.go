package main

import "fmt"

// func main() {
// 	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	resultLen, i, j := m+n-1, m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[resultLen] = nums1[i]
			i--
			resultLen--
		} else {
			nums1[resultLen] = nums2[j]
			j--
			resultLen--
		}
	}
	for j >= 0 {
		nums1[resultLen] = nums2[j]
		j--
		resultLen--
	}
	fmt.Println(nums1)
}
