package tree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root
	res := 0
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		res++
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return res
}
