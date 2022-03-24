package tree

func pathSum(root *TreeNode, targetSum int) (res [][]int) {
	var dfs func(node *TreeNode, path []int, sum int)
	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
	}

	dfs(root, []int{}, 0)
	return
}
