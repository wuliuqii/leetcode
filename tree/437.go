package tree

func pathSum2(root *TreeNode, targetSum int) (res int) {
	preSum := map[int]int{0: 1}
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum += node.Val
		res += preSum[sum-targetSum]
		preSum[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		preSum[sum]--
	}

	dfs(root, 0)
	return
}
