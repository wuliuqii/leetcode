package tree

func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	return dfs(root)
}
