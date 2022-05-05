package tree

func largestValues(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}

	var level []*TreeNode
	level = append(level, root)
	for len(level) != 0 {
		var tmp []*TreeNode
		max := level[0].Val
		for _, node := range level {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		level = tmp
		res = append(res, max)
	}
	return
}
