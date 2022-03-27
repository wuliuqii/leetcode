package tree

func widthOfBinaryTree(root *TreeNode) (res int) {
	type item struct {
		idx int
		*TreeNode
	}
	queue := []item{{0, root}}
	for len(queue) != 0 {
		if l := queue[len(queue)-1].idx - queue[0].idx + 1; l > res {
			res = l
		}
		var tmp []item
		for _, q := range queue {
			if q.Left != nil {
				tmp = append(tmp, item{q.idx * 2, q.Left})
			}
			if q.Right != nil {
				tmp = append(tmp, item{q.idx*2 + 1, q.Right})
			}
		}
		queue = tmp
	}

	return
}
