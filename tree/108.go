package tree

func sortedArrayToBST(nums []int) *TreeNode {
	var buildBST func(begin, end int) *TreeNode
	buildBST = func(begin, end int) *TreeNode {
		if begin == end {
			return nil
		}
		mid := begin + (end-begin)/2
		node := &TreeNode{Val: nums[mid]}
		node.Left = buildBST(begin, mid)
		node.Right = buildBST(mid+1, end)
		return node
	}

	return buildBST(0, len(nums))
}
