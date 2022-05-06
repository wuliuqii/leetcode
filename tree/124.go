package tree

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		res = max(node.Val+left+right, res)
		return max(right+node.Val, left+node.Val)
	}

	dfs(root)
	return res
}
