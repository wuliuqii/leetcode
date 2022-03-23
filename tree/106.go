package tree

func buildTree2(inorder []int, postorder []int) *TreeNode {
	var build func(ii, ij, pi, pj int) *TreeNode
	build = func(ii, ij, pi, pj int) *TreeNode {
		if ii >= ij {
			return nil
		}

		root := &TreeNode{Val: postorder[pj-1]}
		i := ii
		for inorder[i] != root.Val {
			i++
		}
		root.Left = build(ii, i, pi, pi+i-ii)
		root.Right = build(i+1, ij, pi+i-ii, pj-1)
		return root
	}

	return build(0, len(inorder), 0, len(postorder))
}
