package tree

func levelOrder(root *TreeNode) (res [][]int) {
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

	return
}
