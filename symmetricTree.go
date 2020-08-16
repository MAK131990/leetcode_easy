package main

func main() {

}

//TreeNode2 Tree Node
type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func isSymmetric(root *TreeNode2) bool {
	if root == nil {
		return true
	}
	return isSymmetricTrees(root.Left, root.Right)
}

func isSymmetricTrees(left *TreeNode2, right *TreeNode2) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else {
		return left.Val == right.Val && isSymmetricTrees(left.Left, right.Right) && isSymmetricTrees(left.Right, right.Left)
	}
}
