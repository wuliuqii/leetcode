package tree

func levelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		level := make([]int, len(queue))
		var temp []*TreeNode
		for i, node := range queue {
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
			level[i] = node.Val
		}
		queue = temp
		res = append(res, level)
	}

	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}

	return
}
