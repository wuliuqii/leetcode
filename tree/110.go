package tree

func isBalanced(root *TreeNode) bool {
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	if root == nil {
		return true
	}
	if abs(depth(root.Left), depth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
