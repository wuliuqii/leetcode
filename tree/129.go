package tree

func sumNumbers(root *TreeNode) (res int) {
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum = sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += sum
			return
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, 0)
	return
}
