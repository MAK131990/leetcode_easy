package main

// func main() {

// }

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minLeft, minRight := minDepth(root.Left), minDepth(root.Right)
	if minLeft == 0 || minRight == 0 {
		return minLeft + minRight + 1
	}
	return 1 + min(minLeft, minRight)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
