package tree

func zigzagLevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	reverse := false

	for len(queue) != 0 {
		n := len(queue)
		level := make([]int, n)
		var temp []*TreeNode
		for i, node := range queue {
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
			if !reverse {
				level[i] = node.Val
			} else {
				level[n-i-1] = node.Val
			}
		}
		queue = temp
		res = append(res, level)
		reverse = !reverse
	}

	return
}
