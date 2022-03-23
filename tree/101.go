package tree

func isSymmetric(root *TreeNode) bool {
	var same func(left, right *TreeNode) bool
	same = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return same(left.Left, right.Right) && same(left.Right, right.Left)
	}
	if root == nil {
		return true
	}
	return same(root.Left, root.Right)
}
