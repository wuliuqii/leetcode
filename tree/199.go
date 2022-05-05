package tree

func rightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		res = append(res, queue[len(queue)-1].Val)
		var tmp []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}

	return
}
