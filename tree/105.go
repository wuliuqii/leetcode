package tree

func buildTree(preorder []int, inorder []int) *TreeNode {

	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode {

	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]

	// rootVal 在中序遍历中的位置
	var index int = 0

	for i := inStart; i <= inEnd; i++ {

		if rootVal == inorder[i] {
			index = i
			break
		}
	}

	leftSize := index - inStart

	var root = &TreeNode{Val: rootVal}

	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)

	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)

	return root

}
