package tree

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		res = max(res, left+right)
		return max(left, right) + 1
	}

	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
